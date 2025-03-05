package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/inoth/toybox"
	"github.com/inoth/toybox/config"
	"github.com/inoth/toybox/config/file"
	"github.com/inoth/toybox/config/toml"
	"github.com/inoth/toybox/httpserver"
	"github.com/inoth/toybox/metric"
	"github.com/inoth/toybox/profile"
)

var (
	DefaultDir = "config"
)

func newApp(
	conf config.ConfigMate,
	hs *httpserver.GinHttpServer,
	p *metric.Prometheus,
	pprof *profile.Profile,
) *toybox.ToyBox {
	t := toybox.New(
		toybox.WithConfig(conf),
		toybox.WithServer(hs),
		toybox.WithServer(p),
		toybox.WithServer(pprof),
	)
	return t
}

func main() {
	cfg := toml.NewConfiguration(
		config.WithSource(
			file.NewSource(DefaultDir),
		),
	)

restart:
	app := initApp(cfg)
	if err := app.Run(); err != nil && !errors.Is(err, context.Canceled) && !errors.Is(err, toybox.ErrRestart) {
		panic(err)
	} else if errors.Is(err, toybox.ErrRestart) {
		log.Println("restart dbproxy, wait 5s...")
		time.Sleep(time.Second * 5)
		goto restart
	}
}
