package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"

	"lanshan/class11/user/rpc/internal/config"
	"lanshan/class11/user/rpc/internal/server"
	"lanshan/class11/user/rpc/internal/svc"
	"lanshan/class11/user/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "class11/user/rpc/etc/user.sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterSysServer(grpcServer, server.NewSysServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//将consul注册到rpc
	err := consul.RegisterService(c.ListenOn, c.Consul)
	if err != nil {
		fmt.Println("consul error")
		return
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
