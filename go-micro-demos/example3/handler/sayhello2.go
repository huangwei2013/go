package handler

import (
	"context"
	"github.com/lpxxn/gomicrorpc/example3/proto/model"
	"github.com/lpxxn/gomicrorpc/example3/proto/rpcapi"
)

type Say2 struct {}

var _ rpcapi.Say2Handler = (*Say2)(nil)


func (s *Say2) Welcome(ctx context.Context, req *model.SayParam, rsp *model.SayParam) error {
	rsp.Msg = "welcome to here"
	return nil
}

