package main

import (
	"flag"
	"fmt"

	"app/rpc/appserver"
	"app/rpc/internal/config"
	"app/rpc/internal/server"
	"app/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "/Users/jianxinliu/GolandProjects/go-zero-test/rpc/etc/app.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		appserver.RegisterAppServer(grpcServer, server.NewAppServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting appserver server at %s...\n", c.ListenOn)
	s.Start()
}
