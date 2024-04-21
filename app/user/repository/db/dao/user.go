package dao

import (
	"context"
	"todo_list/app/user/repository/db/model"

	"gorm.io/gorm"
)

// 定义的是对数据库的 user model的 CRUD 操作

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) FindUserByUserName(userName string) (r *model.User, err error) {
	err = dao.Model(&model.User{}).
		Where("user_name = ?", userName).Find(&r).Error
	// find 不报： record not found
	// First: SELECT * FROM user WHERE user_name = xxx ORDER BY id LIMIT 1;
	// Find: SELECT * FROM user WHERE user_name = xxx;
	if err != nil {
		return
	}
	return
}

func (dao *UserDao) CreateUser(in *model.User) (err error) {
	return dao.Model(&model.User{}).Create(&in).Error
}
