package main

import (
	"context"
	"os"

	"github.com/inoth/toybox"
	"github.com/inoth/toybox/config"
	"github.com/inoth/toybox/ginsvr"
)

func newApp(conf config.ConfigMate,
	hs *ginsvr.GinHttpServer,
) *toybox.ToyBox {
	t := toybox.New(
		toybox.WithConfig(conf),
		toybox.WithServer(
			hs,
		),
	)
	return t
}

func main() {
	cfgDir := "config"
	if os.Getenv("CONFIG_ENV") == "dev" {
		cfgDir = "../config"
	}
	app := initApp(cfgDir)

	if err := app.Run(); err != nil && err != context.Canceled {
		panic(err)
	}
}
