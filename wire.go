//go:build wireinject

package main

import (
	"github.com/inoth/toybox-layout/internal/biz"
	"github.com/inoth/toybox-layout/internal/data"
	"github.com/inoth/toybox-layout/internal/handler"
	"github.com/inoth/toybox-layout/internal/provider"

	"github.com/google/wire"
	"github.com/inoth/toybox"
	"github.com/inoth/toybox/config"
)

func initApp(conf config.ConfigMate) *toybox.ToyBox {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, handler.ProviderSet, provider.ProviderSet, newApp))
}
