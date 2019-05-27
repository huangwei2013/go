package main

import (
	"context"
	"fmt"
	"github.com/lpxxn/gomicrorpc/example4/common"
	"github.com/lpxxn/gomicrorpc/example4/lib"
	"github.com/lpxxn/gomicrorpc/example4/proto/model"
	"github.com/lpxxn/gomicrorpc/example4/proto/rpcapi"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"io"
	"os"
	"os/signal"
)

func main() {

    //consul
	service := micro.NewService()

	service.Init()
	service.Client().Init(client.Retries(3),
						  client.PoolSize(5))
                          
	demoClient := rpcapi.NewDemoService(common.ServiceName, service.Client())
	userClient := rpcapi.NewUserService(common.ServiceName, service.Client())

	SayHello(demoClient)
	SayMyName(demoClient)
	
    Login(userClient)
    LoginCheck(userClient)
    Logout(userClient)
	
    NotifyTopic(service)
    
	GetStreamValues(demoClient)
	TsBidirectionalStream(demoClient)

	st := make(chan os.Signal)
	signal.Notify(st, os.Interrupt)

	<- st
	fmt.Println("server stopped.....")
}

func SayHello(client rpcapi.DemoService) {
	rsp, err := client.Hello(context.Background(), &model.CommonReq{Action: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

func SayMyName(client rpcapi.DemoService) {
	rsp, err := client.MyName(context.Background(), &model.CommonReq{Action: "Hi server welcome me ?"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

func Login(client rpcapi.UserService) {
	rsp, err := client.Login(context.Background(), &model.CommonReq{Action: "try to login"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

func LoginCheck(client rpcapi.UserService) {
	rsp, err := client.LoginCheck(context.Background(), &model.CommonReq{Action: "try to check login"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

func Logout(client rpcapi.UserService) {
	rsp, err := client.Logout(context.Background(), &model.CommonReq{Action: "try to logout"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

// test stream
func GetStreamValues(client rpcapi.DemoService) {
	rspStream, err := client.Stream(context.Background(), &model.StreamReq{Count: 10})
	if err != nil {
		panic(err)
	}

	idx := 1
	for  {
		rsp, err := rspStream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			//panic(err)
            break            //换成break后，执行正常。但从一般逻辑来说不太合适
		}

		fmt.Printf("test stream get idx %d  data  %v\n", idx, rsp)
		idx++
	}
	// close the stream
	if err := rspStream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
	fmt.Println("Read Value End")
}

func TsBidirectionalStream(client rpcapi.DemoService) {
	rspStream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		panic(err)
	}
	// send
	go func() {
		rspStream.Send(&model.StreamReq{Count: 2})
		rspStream.Send(&model.StreamReq{Count: 5})
		// close the stream
		if err := rspStream.Close(); err != nil {
			fmt.Println("stream close err:", err)
		}
	}()

	idx := 1
	for  {
		rsp, err := rspStream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("test stream get idx %d  data  %v\n", idx, rsp)
		idx++
	}
	fmt.Println("Read Value End")
}

func NotifyTopic(service micro.Service) {
	p := micro.NewPublisher(common.Topic1, service.Client())
	p.Publish(context.TODO(), &model.CommonReq{Action: lib.RandomStr(lib.Random(3, 10))})
}




