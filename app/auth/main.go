package main

import (
	"context"

	auth "github.com/kickastone/trpc-demo/proto/auth"
	"github.com/kickastone/trpc-demo/proto/user"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	auth.RegisterAuthService(s, &AuthServiceImpl{})
	s.Serve()
}

type AuthServiceImpl struct{}

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginRequest) (rsp *auth.LoginResponse, err error) {
	rsp = &auth.LoginResponse{}
	uReq := &user.GetAccountByUserNameRequest{
		Metadata: req.GetMetadata(),
		Username: req.GetUsername(),
	}
	uRsp, err := user.NewUserClientProxy().GetAccountByUserName(ctx, uReq)
	if err != nil {
		log.ErrorContextf(ctx, "get account by user name failed, err: %v", err)
		return nil, err
	}

	if uRsp.GetErrCode() != 0 {
		rsp.ErrCode = uRsp.GetErrCode()
		rsp.ErrMsg = uRsp.GetErrMsg()
		return
	}

	log.Infof("password is correct, user: %v, username: %v, password: %v, password hash: %v", req.GetUsername(), uRsp.Data.UserInfo.GetUsername(), req.GetPassword(), uRsp.Data.UserInfo.GetPasswordHash())

	if uRsp.Data.UserInfo.GetPasswordHash() != req.GetPassword() {
		rsp.ErrCode, rsp.ErrMsg = -1, "password is incorrect"
	}

	return rsp, nil
}
