package user

import (
	"context"

	"github.com/inoth/toybox-layout/internal/biz"
	"github.com/inoth/toybox-layout/internal/data/user/model"

	database "github.com/inoth/toybox/component/database/sqlite"
	"github.com/inoth/toybox/logger"
	"github.com/inoth/toybox/util/convert"
	"gorm.io/gorm"
)

type userRepo struct {
	log logger.Logger
	db  *gorm.DB
}

func NewUserRepo(log logger.Logger, db *database.SqliteComponent) biz.UserRepo {
	user := db.GetDatabase("toybox-layout", &model.UserInfo{})
	return &userRepo{
		log: log,
		db:  user,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.UserInfo) (uint, error) {

	r.log.Log(ctx, logger.LevelInfo, "CreateUser")

	add, err := convert.ConvertByUnsafe[biz.UserInfo, model.UserInfo](user)
	if err != nil {
		return 0, err
	}
	if err := r.db.Create(add).Error; err != nil {
		return 0, err
	}
	return add.Id, nil
}

func (r *userRepo) GetUserById(ctx context.Context, id uint) (*biz.UserInfo, error) {

	r.log.Log(ctx, logger.LevelInfo, "GetUserById")

	var user model.UserInfo
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return convert.ConvertByUnsafe[model.UserInfo, biz.UserInfo](&user)
}

func (r *userRepo) GetAllUsers(ctx context.Context) ([]*biz.UserInfo, error) {

	r.log.Log(ctx, logger.LevelInfo, "GetAllUsers")

	var users []model.UserInfo
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	res := make([]*biz.UserInfo, 0, len(users))
	for _, user := range users {
		r, err := convert.ConvertByUnsafe[model.UserInfo, biz.UserInfo](&user)
		if err != nil {
			continue
		}
		res = append(res, r)
	}
	return res, nil
}
