package main

import (
	"fmt"
	"time"

	"todo_list/app/gateway/router"
	"todo_list/app/gateway/rpc"
	"todo_list/app/user/repository/cache"
	"todo_list/config"
	log "todo_list/pkg/logger"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/web"

	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	rpc.InitRPC()
	cache.InitCache()
	log.InitLog()
	// etcd 注册
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// new 一个微服务实例，使用gin暴露http接口并注册到etcd
	webService := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		// 将服务调用实例使用gin处理
		web.Handler(router.NewRouter()),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}))

	_ = webService.Init()
	_ = webService.Run()

}
