package handler

import (
	"context"
	"github.com/lpxxn/gomicrorpc/example4/proto/model"
	"github.com/lpxxn/gomicrorpc/example4/proto/rpcapi"
)

type User struct {}

var _ rpcapi.UserHandler = (*User)(nil)


func (s *User) Login(ctx context.Context, req *model.CommonReq, rsp *model.CommonRsp) error {
	rsp.Msg = "welcome to here"
	return nil
}

func (s *User) Logout(ctx context.Context, req *model.CommonReq, rsp *model.CommonRsp) error {
	rsp.Msg = "welcome to here"
	return nil
}

func (s *User) LoginCheck(ctx context.Context, req *model.CommonReq, rsp *model.CommonRsp) error {
	rsp.Msg = "welcome to here"
	return nil
}
