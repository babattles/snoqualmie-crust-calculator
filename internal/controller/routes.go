package controller

import (
	"io/fs"
	"net/http"

	"github.com/babattles/snoqualmie-crust-calculator/config"
	"github.com/babattles/snoqualmie-crust-calculator/web"
	"github.com/labstack/echo/v4"
)

func BindRoutes(e *echo.Echo, c Controller, cfg config.Config) error {
	staticFS, err := fs.Sub(web.FS, "static")
	if err != nil {
		return err
	}
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", http.FileServer(http.FS(staticFS)))))
	e.GET("/", c.GetIndex)

	api := e.Group("/api", APIKeyMiddleware(cfg.APIKey))
	api.GET("/health", c.GetHealthCheck)
	api.GET("/crusts", c.GetCrusts)
	return nil
}
