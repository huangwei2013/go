package main

import (
	"context"
	"fmt"
	"github.com/lpxxn/gomicrorpc/example3/common"
	"github.com/lpxxn/gomicrorpc/example3/lib"
	"github.com/lpxxn/gomicrorpc/example3/proto/model"
	"github.com/lpxxn/gomicrorpc/example3/proto/rpcapi"
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
                          
	sayClent := rpcapi.NewSayService(common.ServiceName, service.Client())
	say2Clent := rpcapi.NewSay2Service(common.ServiceName, service.Client())

	SayHello(sayClent)
	SayMyName(sayClent)
	SayWelcome(say2Clent)
	NotifyTopic(service)
	GetStreamValues(sayClent)
	TsBidirectionalStream(sayClent)

	st := make(chan os.Signal)
	signal.Notify(st, os.Interrupt)

	<- st
	fmt.Println("server stopped.....")
}

func SayHello(client rpcapi.SayService) {
	rsp, err := client.Hello(context.Background(), &model.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

func SayMyName(client rpcapi.SayService) {
	rsp, err := client.MyName(context.Background(), &model.SayParam{Msg: "Hi server welcome me ?"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

func SayWelcome(client rpcapi.Say2Service) {
	rsp, err := client.MyName(context.Background(), &model.SayParam{Msg: "Hi server welcome me ?"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

// test stream
func GetStreamValues(client rpcapi.SayService) {
	rspStream, err := client.Stream(context.Background(), &model.SRequest{Count: 10})
	if err != nil {
		panic(err)
	}

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
	// close the stream
	if err := rspStream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
	fmt.Println("Read Value End")
}

func TsBidirectionalStream(client rpcapi.SayService) {
	rspStream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		panic(err)
	}
	// send
	go func() {
		rspStream.Send(&model.SRequest{Count: 2})
		rspStream.Send(&model.SRequest{Count: 5})
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
	p.Publish(context.TODO(), &model.SayParam{Msg: lib.RandomStr(lib.Random(3, 10))})
}




