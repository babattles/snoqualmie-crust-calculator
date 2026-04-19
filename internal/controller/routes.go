package controller

import (
	"github.com/babattles/snoqualmie-crust-calculator/config"
	"github.com/labstack/echo/v4"
)

func BindRoutes(e *echo.Echo, c Controller, cfg config.Config) {
	api := e.Group("/api", APIKeyMiddleware(cfg.APIKey))
	api.GET("/health", c.GetHealthCheck)
	api.GET("/crusts", c.GetCrusts)
}
