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

	// 启动一些脚本
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

	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterTaskServiceHandler(microService.Server(), service.GetTaskSrv())
	// 启动微服务
	_ = microService.Run()
}

func loadingScript() {
	ctx := context.Background()
	go script.TaskCreateSync(ctx)
}
