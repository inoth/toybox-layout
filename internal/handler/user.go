package handler

import (
	"strconv"

	"github.com/inoth/toybox-layout/internal/biz"

	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox/httpserver"
	"github.com/inoth/toybox/httpserver/res"
	"github.com/inoth/toybox/logger"
	"github.com/inoth/toybox/util"
)

type UserInfoHandler struct {
	log logger.Logger
	svr *biz.UserUsecase
}

func NewUserInfoHandler(log logger.Logger, svr *biz.UserUsecase) *UserInfoHandler {
	return &UserInfoHandler{
		log: log,
		svr: svr,
	}
}

func (uc *UserInfoHandler) Prefix() string {
	return "/api/user"
}

func (uc *UserInfoHandler) Middlewares() []gin.HandlerFunc {
	return nil
}

func (uc *UserInfoHandler) Routers() []*httpserver.Router {
	return []*httpserver.Router{
		httpserver.NewRouter("GET", "", uc.GetUser),
		httpserver.NewRouter("GET", "/:id", uc.GetUserById),
		httpserver.NewRouter("GET", "/add", uc.CreateUserRandom),
	}
}

func (uc *UserInfoHandler) GetUserById(c *gin.Context) {

	uc.log.Log(c, logger.LevelInfo, "GetUserById()")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.ErrParams(c, err.Error())
		return
	}
	userInfo, err := uc.svr.GetUserById(c, uint(id))
	if err != nil {
		res.ErrParams(c, err.Error())
		return
	}
	res.Ok(c, "ok", userInfo)
}

func (uc *UserInfoHandler) CreateUserRandom(c *gin.Context) {

	uc.log.Log(c, logger.LevelInfo, "CreateUserRandom()")

	userId, err := uc.svr.CreateUser(c, &biz.UserInfo{Name: util.UUID(), Email: util.UUID(10) + "@example.com"})
	if err != nil {
		res.Failed(c, err.Error())
		return
	}
	res.Ok(c, "ok", userId)
}

func (uc *UserInfoHandler) GetUser(c *gin.Context) {

	uc.log.Log(c, logger.LevelInfo, "GetUser()")

	users, err := uc.svr.GetAllUsers(c)
	if err != nil {
		res.Failed(c, err.Error())
		return
	}
	res.Ok(c, "ok", users)
}
