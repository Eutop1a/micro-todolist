package rpc

import (
	"context"
	"todo_list/idl/pb"
	"todo_list/pkg/e"
)

func UserLogin(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserLogin(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}
func UserRegister(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserRegister(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}
