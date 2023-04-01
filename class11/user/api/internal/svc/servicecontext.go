package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"lanshan/class11/user/api/internal/config"
	"lanshan/class11/user/model"
	"lanshan/class11/user/rpc/sys"
)

type ServiceContext struct {
	Config       config.Config
	UserModel    model.UserModel
	SysRpcClient sys.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		UserModel:    model.NewUserModel(conn, c.CacheRedis),
		SysRpcClient: sys.NewSys(zrpc.MustNewClient(c.SysRpcClientConf)),
	}
}
