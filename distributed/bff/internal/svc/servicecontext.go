package svc

import (
	"bff/internal/config"
	"user/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserSrv user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserSrv: user.NewUserZrpcClient(zrpc.MustNewClient(c.UserSrv)),
	}
}
