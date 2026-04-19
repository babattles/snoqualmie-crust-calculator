package controller

import (
	"errors"
	"net/http"

	"github.com/babattles/snoqualmie-crust-calculator/internal/entity"
	"github.com/babattles/snoqualmie-crust-calculator/internal/usecase"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	usecase *usecase.Usecase
}

func NewController(u *usecase.Usecase) Controller {
	return Controller{usecase: u}
}

func (ctrl Controller) GetHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func (ctrl Controller) GetCrusts(c echo.Context) error {
	mountain := c.QueryParam("mountain")
	if mountain == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "mountain query param is required")
	}

	resp, err := ctrl.usecase.GetCrustsForMountain(c.Request().Context(), entity.MountainType(mountain))
	if err != nil {
		if errors.Is(err, usecase.ErrUnknownMountain) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
