package greeter

import (
	"fmt"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/urfave/cli"

	"kratos_template/infrastructure"
	"kratos_template/internal/application/greeter"
	"kratos_template/internal/conf"
	domainSvc "kratos_template/internal/domain/greeter/service"
	"kratos_template/internal/server"
	"kratos_template/internal/service"
)

var Command = cli.Command{
	Name:  "greeter-service",
	Usage: "start greeter service",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "configs/config.yaml",
			Usage: "path to config file",
		},
	},
	Action: func(c *cli.Context) error {
		return start(c.String("config"))
	},
}

func start(configPath string) error {
	cfg, err := conf.Load(configPath)
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}

	data, cleanup, err := infrastructure.NewData(cfg.Data.Database)
	if err != nil {
		return fmt.Errorf("initializing data: %w", err)
	}
	defer cleanup()

	// Domain
	greeterRepo := infrastructure.NewGreeterRepo(data)
	greeterDomainSvc := domainSvc.NewGreeterService(greeterRepo)

	// Application
	appSvc := greeter.NewAppService(greeterDomainSvc)

	// Service (proto interface implementation)
	greeterSvc := service.NewGreeterService(appSvc)

	// Servers
	grpcSrv := server.NewGRPCServer(&cfg.Server, greeterSvc)
	httpSrv := server.NewHTTPServer(&cfg.Server, greeterSvc)

	// Kratos app
	logger := log.DefaultLogger
	app := kratos.New(
		kratos.Name("greeter"),
		kratos.Server(grpcSrv, httpSrv),
		kratos.Logger(logger),
	)

	return app.Run()
}
