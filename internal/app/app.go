package app

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/babattles/snoqualmie-crust-calculator/config"
	"github.com/babattles/snoqualmie-crust-calculator/internal/controller"
	"github.com/babattles/snoqualmie-crust-calculator/internal/repo/webapi/snowobs"
	"github.com/babattles/snoqualmie-crust-calculator/internal/usecase"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Providers = fx.Options(
	fx.Provide(config.New),
	fx.Provide(NewServer),
	fx.Provide(NewSnowobsClient),
	usecase.Modules,
	controller.Modules,
)

var Modules = fx.Options(
	Providers,
	fx.Invoke(RegisterHooks),
)

func NewServer() *echo.Echo {
	return echo.New()
}

func NewSnowobsClient(cfg config.Config) *snowobs.Client {
	return snowobs.New(snowobs.Config{
		Token:  cfg.SnowobsToken,
		Source: cfg.SnowobsSource,
		BaseURL: cfg.SnowobsBaseURL,
	})
}

func RegisterHooks(lifecycle fx.Lifecycle, e *echo.Echo, shutdowner fx.Shutdowner) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Printf("echo server error: %v", err)
					_ = shutdowner.Shutdown()
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
