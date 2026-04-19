package controller

import "github.com/labstack/echo/v4"

func BindRoutes(e *echo.Echo, c Controller) {
	e.GET("/health", c.GetHealthCheck)
	e.GET("/crusts", c.GetCrusts)
}
