package controller

import (
	"net/http"
	"time"

	"github.com/babattles/snoqualmie-crust-calculator/internal/entity"
	"github.com/babattles/snoqualmie-crust-calculator/internal/pkg/crust"
	"github.com/labstack/echo/v4"
)

type indexLayer struct {
	ElevationFt    int
	SunCrust       crust.CrustType
	MeltCrust      crust.CrustType
	SunCrustLabel  string
	MeltCrustLabel string
}

type indexView struct {
	Mountain  entity.MountainType
	FetchedAt string
	Layers    []indexLayer
	Error     string
}

var crustLabels = map[crust.CrustType]string{
	crust.CrustSun:      "Likely",
	crust.CrustSunMaybe: "Possible",
	crust.CrustRain:     "Rain crust",
	crust.CrustMelt:     "Melt crust",
	crust.CrustNone:     "None",
}

func crustLabel(c crust.CrustType) string {
	if label, ok := crustLabels[c]; ok {
		return label
	}
	return string(c)
}

func (ctrl Controller) GetIndex(c echo.Context) error {
	mountain := entity.MountainTypeAlpental
	resp, err := ctrl.usecase.GetCrustsForMountain(c.Request().Context(), mountain)
	if err != nil {
		return c.Render(http.StatusOK, "index", indexView{
			Mountain: mountain,
			Error:    "Could not load weather data right now. Try again in a minute.",
		})
	}

	layers := make([]indexLayer, len(resp.Layers))
	for i, l := range resp.Layers {
		layers[i] = indexLayer{
			ElevationFt:    l.ElevationFt,
			SunCrust:       l.SunCrust,
			MeltCrust:      l.MeltCrust,
			SunCrustLabel:  crustLabel(l.SunCrust),
			MeltCrustLabel: crustLabel(l.MeltCrust),
		}
	}

	return c.Render(http.StatusOK, "index", indexView{
		Mountain:  resp.Mountain,
		FetchedAt: time.Now().Format("Mon Jan 2 3:04pm MST"),
		Layers:    layers,
	})
}
