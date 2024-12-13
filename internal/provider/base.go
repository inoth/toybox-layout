package provider

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewConfigMate,
	NewHttpServer,
)