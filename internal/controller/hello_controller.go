package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox-layout/internal/service"
	"github.com/inoth/toybox/ginsvr"
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

func (h *HelloController) Routers() []ginsvr.Router {
	return []ginsvr.Router{
		{Method: "GET", Path: "/sayhi/:name", Handle: []gin.HandlerFunc{h.SayHi}},
	}
}

func (h *HelloController) SayHi(c *gin.Context) {
	name := c.Param("name")
	c.String(200, h.svr.SayHi(name))
}
