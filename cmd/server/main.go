package main

import (
	"github.com/babattles/snoqualmie-crust-calculator/internal/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(app.Modules).Run()
}
