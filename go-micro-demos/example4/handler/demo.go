package handler

import (
	"context"
	"fmt"
	"github.com/lpxxn/gomicrorpc/example4/lib"
	"github.com/lpxxn/gomicrorpc/example4/proto/model"
	"github.com/lpxxn/gomicrorpc/example4/proto/rpcapi"
	"io"
	"time"
)

type Demo struct {}

var _ rpcapi.DemoHandler = (*Demo)(nil)

func (s *Demo) Hello(ctx context.Context, req *model.CommonReq, rsp *model.CommonRsp) error {
	fmt.Println("received", req.Action)
	rsp.Data = make(map[string]*model.Pair)
	rsp.Data["name"] = &model.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello world"
	rsp.Code = 0

	return nil
}

func (s *Demo) MyName(ctx context.Context, req *model.CommonReq, rsp *model.CommonRsp) error {
	rsp.Msg = "lp"
	return nil
}

/*
 模拟得到一些数据
 */
func (s *Demo) Stream(ctx context.Context, req *model.StreamReq, stream rpcapi.Demo_StreamStream) error {

	for i := 0; i < int(10); i++ {
		rsp := &model.StreamRsp{}
        rsp.Code = 0
		for j := lib.Random(3, 5); j < 10; j++ {
            rsp.Data = append(rsp.Data, lib.RandomStr(lib.Random(3, 10)))
		}
		if err := stream.Send(rsp); err != nil {
			return err
		}
		// 模拟处理过程
		time.Sleep(time.Microsecond * 50)
	}
	return nil
}

/*
 模拟数据
 */
func (s *Demo) BidirectionalStream(ctx context.Context, stream rpcapi.Demo_BidirectionalStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if req == nil {
			return err
		}
		for i := int64(0); i < 10; i++ {
			if err := stream.Send(&model.StreamRsp{Code : 0, Data: []string {lib.RandomStr(lib.Random(3, 6))}}); err != nil {
                return err
			}
		}
	}
	return nil
}

