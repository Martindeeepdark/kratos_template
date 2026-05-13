package server

import (
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "kratos_template/api/helloworld/v1"
	"kratos_template/internal/conf"
	"kratos_template/internal/service"
)

func NewHTTPServer(c *conf.ServerConfig, greeterSvc *service.GreeterService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.HTTP.Network != "" {
		opts = append(opts, http.Network(c.HTTP.Network))
	}
	if c.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.HTTP.Addr))
	}
	if c.HTTP.Timeout > 0 {
		opts = append(opts, http.Timeout(c.HTTP.Timeout))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeterSvc)
	return srv
}
