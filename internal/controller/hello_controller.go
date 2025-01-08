package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox-layout/internal/service"
	"github.com/inoth/toybox/ginserver"
	"github.com/inoth/toybox/ginserver/res"
)

type HelloController struct {
	svr *service.HelloService
}

func NewHelloController(svr *service.HelloService) *HelloController {
	return &HelloController{
		svr: svr,
	}
}

func (h *HelloController) Prefix() string { return "/api" }

func (h *HelloController) Middlewares() []gin.HandlerFunc { return nil }

func (h *HelloController) Routers() []ginserver.Router {
	return []ginserver.Router{
		{Method: "GET", Path: "/sayhi/:name", Handle: []gin.HandlerFunc{h.SayHi}},
	}
}

func (h *HelloController) SayHi(c *gin.Context) {
	name := c.Param("name")
	res.Ok(c, h.svr.SayHi(name))
}
