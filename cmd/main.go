package main

import (
	"contest/internal/app"
	"contest/internal/config"
)

func main() {
	cfg := config.MustLoadAppConfig()
	application := app.New(cfg)
	application.MustRun()
}
