package main

import (
	"fmt"
	"time"
	"todo_list/app/gateway/router"
	"todo_list/app/gateway/rpc"
	"todo_list/config"

	"go-micro.dev/v4/web"

	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	rpc.InitRPC()
	// etcd 注册
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))

	// new 一个微服务实例
	webService := web.NewService(
		web.Name("httpService"),
		web.Address("localhost:4000"),
		web.Registry(etcdReg),
		web.Handler(router.NewRouter()),
		web.RegisterTTL(time.Second*30),
		web.Metadata(map[string]string{"protocol": "http"}))

	webService.Init()

	_ = webService.Run()

}
