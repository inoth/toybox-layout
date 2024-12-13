package provider

import "github.com/inoth/toybox/config"

func NewConfigMate(dir string) config.ConfigMate {
	return config.NewConfig(
		config.WithConfigDir(dir),
	)
}
