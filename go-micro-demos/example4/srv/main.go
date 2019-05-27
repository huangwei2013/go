package main

import (
	"github.com/lpxxn/gomicrorpc/example4/common"
	"github.com/lpxxn/gomicrorpc/example4/handler"
	"github.com/lpxxn/gomicrorpc/example4/proto/rpcapi"
	"github.com/lpxxn/gomicrorpc/example4/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"time"
)


func main() {

	// 初始化服务 by consuls
	service := micro.NewService(
		micro.Name(common.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
	)

	service.Init()
    
	// 注册 Handler
	rpcapi.RegisterDemoHandler(service.Server(), new(handler.Demo))
	rpcapi.RegisterUserHandler(service.Server(), new(handler.User))


	// Register Subscribers
	if err := server.Subscribe(server.NewSubscriber(common.Topic1, subscriber.Handler)); err != nil {
		panic(err)
	}

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
