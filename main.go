package main

import (
	"context"

	"github.com/inoth/toybox/config"
)

func main() {
	app := initApp(config.CfgBasic{
		Remote:   false,
		CfgDir:   "config",
		FileType: "toml",
		Env:      "",
	})

	if err := app.Run(); err != nil && err != context.Canceled {
		panic(err)
	}
}
