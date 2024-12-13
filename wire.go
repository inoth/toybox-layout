//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/inoth/toybox"
	"github.com/inoth/toybox-layout/internal/controller"
	"github.com/inoth/toybox-layout/internal/provider"
	"github.com/inoth/toybox-layout/internal/service"
)

func initApp(dir string) *toybox.ToyBox {
	panic(wire.Build(service.ProviderSet, controller.ProviderSet, provider.ProviderSet, newApp))
}
