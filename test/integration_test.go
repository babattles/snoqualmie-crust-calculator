package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/babattles/snoqualmie-crust-calculator/internal/app"
	"github.com/babattles/snoqualmie-crust-calculator/internal/entity"
	"github.com/babattles/snoqualmie-crust-calculator/internal/pkg/crust"
	"github.com/babattles/snoqualmie-crust-calculator/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

const testAPIKey = "test-api-key"

func newMockSnowobsServer(t *testing.T) *httptest.Server {
	t.Helper()
	body, err := os.ReadFile(filepath.Join("testdata", "snowobs_response.json"))
	require.NoError(t, err)

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(body)
	}))
}

func buildTestApp(t *testing.T, snowobsBaseURL string) *echo.Echo {
	t.Helper()
	t.Setenv("SNOWOBS_BASE_URL", snowobsBaseURL)
	t.Setenv("SNOWOBS_TOKEN", "test-token")
	t.Setenv("SNOWOBS_SOURCE", "nwac")
	t.Setenv("API_KEY", testAPIKey)

	var e *echo.Echo
	fxApp := fx.New(
		app.Providers,
		fx.Populate(&e),
		fx.NopLogger,
	)
	require.NoError(t, fxApp.Err())
	return e
}

func doGet(t *testing.T, url, apiKey string) *http.Response {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	if apiKey != "" {
		req.Header.Set("X-API-Key", apiKey)
	}
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	return resp
}

func TestIntegration_Endpoints(t *testing.T) {
	snowobsMock := newMockSnowobsServer(t)
	defer snowobsMock.Close()

	e := buildTestApp(t, snowobsMock.URL)
	ts := httptest.NewServer(e)
	defer ts.Close()

	t.Run("GET /api/health returns 200 ok", func(t *testing.T) {
		resp := doGet(t, ts.URL+"/api/health", testAPIKey)
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "ok", string(body))
	})

	t.Run("GET /api/crusts?mountain=Alpental returns sorted crust layers", func(t *testing.T) {
		resp := doGet(t, ts.URL+"/api/crusts?mountain=Alpental", testAPIKey)
		defer resp.Body.Close()
		require.Equal(t, http.StatusOK, resp.StatusCode)

		var result usecase.MountainCrusts
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&result))

		assert.Equal(t, entity.MountainTypeAlpental, result.Mountain)
		require.Len(t, result.Layers, 3)

		assert.Equal(t, 3100, result.Layers[0].ElevationFt)
		assert.Equal(t, crust.CrustNone, result.Layers[0].SunCrust)
		assert.Equal(t, crust.CrustMelt, result.Layers[0].MeltCrust)

		assert.Equal(t, 4350, result.Layers[1].ElevationFt)
		assert.Equal(t, crust.CrustSun, result.Layers[1].SunCrust)
		assert.Equal(t, crust.CrustMelt, result.Layers[1].MeltCrust)

		assert.Equal(t, 5470, result.Layers[2].ElevationFt)
		assert.Equal(t, crust.CrustSunMaybe, result.Layers[2].SunCrust)
		assert.Equal(t, crust.CrustMelt, result.Layers[2].MeltCrust)
	})

	t.Run("GET /api/crusts without mountain returns 400", func(t *testing.T) {
		resp := doGet(t, ts.URL+"/api/crusts", testAPIKey)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("GET /api/crusts with unknown mountain returns 400", func(t *testing.T) {
		resp := doGet(t, ts.URL+"/api/crusts?mountain=nonexistent", testAPIKey)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("GET /api/crusts without API key returns 401", func(t *testing.T) {
		resp := doGet(t, ts.URL+"/api/crusts?mountain=Alpental", "")
		defer resp.Body.Close()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("GET /api/crusts with wrong API key returns 401", func(t *testing.T) {
		resp := doGet(t, ts.URL+"/api/crusts?mountain=Alpental", "wrong-key")
		defer resp.Body.Close()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("GET /api/health without API key returns 401", func(t *testing.T) {
		resp := doGet(t, ts.URL+"/api/health", "")
		defer resp.Body.Close()
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}
