package rpc

import (
	"todo_list/app/gateway/wrappers"
	"todo_list/idl/pb"

	"go-micro.dev/v4"
)

var (
	UserService pb.UserService
	TaskService pb.TaskService
)

func InitRPC() {
	// 用户

	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	// 用户服务调用实例
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())

	// task
	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
	)
	taskService := pb.NewTaskService("rpcTaskService", taskMicroService.Client())

	UserService = userService
	TaskService = taskService
}
