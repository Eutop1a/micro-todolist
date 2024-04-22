package main

import (
	"context"
	"fmt"
	"todo_list/app/task/repository/db/dao"
	"todo_list/app/task/repository/mq"
	"todo_list/app/task/script"
	"todo_list/app/task/service"
	"todo_list/config"
	"todo_list/idl/pb"

	"go-micro.dev/v4"

	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitDB()
	mq.InitRabbitMQ()
	loadingScript()
	// etcd 注册
	etcdReg := registry.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// new 一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address(config.TaskServiceAddress),
		micro.Registry(etcdReg),
	)

	microService.Init()
	_ = pb.RegisterTaskServiceHandler(microService.Server(), service.GetTaskSrv())

	_ = microService.Run()
}

func loadingScript() {
	ctx := context.Background()
	go script.TaskCreateSync(ctx)
}
