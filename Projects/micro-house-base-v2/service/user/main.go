package main

import (
	"micro-house-base-v2/service/user/handler"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	"micro-house-base-v2/service/user/model"
	user "micro-house-base-v2/service/user/proto/user"

	"github.com/go-micro/plugins/v2/registry/consul"
)

func main() {
	// 初始化 MySQL 连接池
	model.InitDb()
	// 初始化 redis 连接池
	model.InitRedis()

	// 初始化 Consul
	consulReg := consul.NewRegistry()

	// New Service  -- 指定 consul
	service := micro.NewService(
		micro.Address("192.168.6.108:12342"), // 指定固定端口
		micro.Name("go.micro.srv.user"),
		micro.Registry(consulReg), // 注册服务
		micro.Version("latest"),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
