package provider

import (
	"github.com/inoth/toybox-layout/internal/controller"
	"github.com/inoth/toybox/ginsvr"
)

func NewHttpServer(h *controller.HelloController) *ginsvr.GinHttpServer {
	return ginsvr.New(
		ginsvr.WithHandlers(h),
	)
}
