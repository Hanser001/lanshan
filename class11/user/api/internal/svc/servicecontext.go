package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lanshan/class11/user/api/internal/config"
	"lanshan/class11/user/rpc/sys"
)

type ServiceContext struct {
	Config       config.Config
	SysRpcClient sys.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		SysRpcClient: sys.NewSys(zrpc.MustNewClient(c.SysRpcClientConf)),
	}
}
