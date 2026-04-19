package controller

import (
	"crypto/subtle"
	"net/http"

	"github.com/babattles/snoqualmie-crust-calculator/config"
	"github.com/labstack/echo/v4"
)

const apiKeyHeader = "X-API-Key"

func APIKeyMiddleware(apiKey string) echo.MiddlewareFunc {
	expected := []byte(apiKey)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if len(expected) == 0 {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			provided := []byte(c.Request().Header.Get(apiKeyHeader))
			if subtle.ConstantTimeCompare(provided, expected) != 1 {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			return next(c)
		}
	}
}

func RegisterMiddleware(e *echo.Echo, cfg config.Config) {
	e.Use(APIKeyMiddleware(cfg.APIKey))
}
