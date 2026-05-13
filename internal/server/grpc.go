package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	v1 "kratos_template/api/helloworld/v1"
	"kratos_template/internal/conf"
	"kratos_template/internal/service"
)

func NewGRPCServer(c *conf.ServerConfig, greeterSvc *service.GreeterService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.GRPC.Network != "" {
		opts = append(opts, grpc.Network(c.GRPC.Network))
	}
	if c.GRPC.Addr != "" {
		opts = append(opts, grpc.Address(c.GRPC.Addr))
	}
	if c.GRPC.Timeout > 0 {
		opts = append(opts, grpc.Timeout(c.GRPC.Timeout))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeterSvc)
	return srv
}
