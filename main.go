package main

import (
	"flag"
	"fmt"

	"github.com/wuyazi/grpc_user_domain/grpc_user_domain"
	"github.com/wuyazi/grpc_user_domain/internal/config"
	"github.com/wuyazi/grpc_user_domain/internal/server"
	"github.com/wuyazi/grpc_user_domain/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewGrpcUserDomainServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		grpc_user_domain.RegisterUserCommendServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
