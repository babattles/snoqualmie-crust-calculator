package controller

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(NewController),
	fx.Provide(NewRenderer),
	fx.Invoke(attachRenderer),
	fx.Invoke(BindRoutes),
)

func attachRenderer(e *echo.Echo, r *Renderer) {
	e.Renderer = r
}
