package biz

import (
	"context"

	"github.com/inoth/toybox/logger"
)

type UserInfo struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *UserInfo) (uint, error)
	GetUserById(ctx context.Context, id uint) (*UserInfo, error)
	GetAllUsers(ctx context.Context) ([]*UserInfo, error)
}

type UserUsecase struct {
	log      logger.Logger
	userRepo UserRepo
}

func NewUserUsecase(log logger.Logger, userRepo UserRepo) *UserUsecase {
	return &UserUsecase{
		log:      log,
		userRepo: userRepo,
	}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *UserInfo) (uint, error) {
	uc.log.Log(ctx, logger.LevelInfo, "CreateUser")
	return uc.userRepo.CreateUser(ctx, user)
}

func (uc *UserUsecase) GetUserById(ctx context.Context, id uint) (*UserInfo, error) {
	uc.log.Log(ctx, logger.LevelInfo, "GetUserById")
	return uc.userRepo.GetUserById(ctx, id)
}

func (uc *UserUsecase) GetAllUsers(ctx context.Context) ([]*UserInfo, error) {
	uc.log.Log(ctx, logger.LevelInfo, "GetAllUsers")
	return uc.userRepo.GetAllUsers(ctx)
}
