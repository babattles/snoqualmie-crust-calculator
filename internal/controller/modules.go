package controller

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewController),
	fx.Invoke(RegisterMiddleware),
	fx.Invoke(BindRoutes),
)
