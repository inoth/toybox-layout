package provider

import (
	"github.com/inoth/toybox-layout/internal/controller"
	"github.com/inoth/toybox/ginserver"
	"github.com/inoth/toybox/ginserver/middleware"
)

func NewHttpServer(h *controller.HelloController) *ginserver.GinHttpServer {
	return ginserver.NewHttp(
		ginserver.WithMiddleware(
			middleware.SetTraceId(),
		),
		ginserver.WithHandlers(h),
	)
}
