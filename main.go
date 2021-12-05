package main

import (
	"github.com/Heilartin/bot_support/app"
	"github.com/Heilartin/bot_support/config"
	"github.com/Heilartin/bot_support/logger"
	"github.com/k0kubun/pp"
)

func main() {
	pp.Println()
	cfg := config.NewConfig()
	ll := logger.NewLogger(cfg.ProductionStart)
	ap, err := app.New(cfg, ll)
	if err != nil {
		ll.Fatal(err)
	}
	ap.Run()
}
