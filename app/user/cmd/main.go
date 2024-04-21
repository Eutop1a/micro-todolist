package main

import (
	"fmt"
	"todo_list/app/user/repository/db/dao"
	"todo_list/app/user/service"
	"todo_list/config"
	"todo_list/idl/pb"

	"go-micro.dev/v4"

	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitDB()

	// etcd 注册
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))

	// new 一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg))

	microService.Init()
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())

	_ = microService.Run()
}
