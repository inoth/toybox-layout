//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/inoth/toybox"
	"github.com/inoth/toybox-layout/internal/controller"
	"github.com/inoth/toybox-layout/internal/register"
	"github.com/inoth/toybox-layout/internal/service"
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

func initApp(cfg config.CfgBasic) *toybox.ToyBox {
	panic(wire.Build(config.NewConfig, service.ProviderSet, controller.ProviderSet, register.ProviderSet, newApp))
}
