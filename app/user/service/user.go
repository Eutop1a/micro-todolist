package service

import (
	"context"
	"errors"
	"sync"
	"todo_list/app/user/repository/db/dao"
	"todo_list/app/user/repository/db/model"
	"todo_list/idl/pb"
	"todo_list/pkg/e"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

// GetUserSrv 懒汉式的单例模式 lazy-loading
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

// GetUserSrvHungry 饿汉式
func GetUserSrvHungry() *UserSrv {
	if UserSrvIns == nil {
		return new(UserSrv)
	}
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = e.Success
	// 查看有没有这个人
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID == 0 {
		err = errors.New("用户不存在")
		resp.Code = e.Error
		return
	}
	if !user.CheckPassword(req.Password) {
		err = errors.New("用户密码错误")
		resp.Code = e.Error
		return
	}
	resp.UserDetail = BuildUser(user)
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserResponse) (err error) {
	resp.Code = e.Success
	if req.Password != req.PasswordConfirm {
		err = errors.New("两次密码不一致")
		resp.Code = e.Error
		return
	}
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		return
	}
	if user.ID > 0 {
		err = errors.New("用户名已存在")
		resp.Code = e.Error
		return
	}
	user = &model.User{
		UserName: req.UserName,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = e.Error
		return
	}

	if err = dao.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = e.Error
		return
	}
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
}
