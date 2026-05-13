package service

import (
	"context"

	"github.com/Martindeeepdark/go-common/errorx"
	"github.com/Martindeeepdark/go-common/logs"
	"github.com/go-kratos/kratos/v2/errors"

	v1 "kratos_template/api/helloworld/v1"
	"kratos_template/internal/application/greeter"
)

type GreeterService struct {
	v1.UnimplementedGreeterServer

	app *greeter.AppService
}

func NewGreeterService(app *greeter.AppService) *GreeterService {
	return &GreeterService{app: app}
}

func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	resp, err := s.app.CreateGreeter(ctx, &greeter.CreateGreeterRequest{
		Hello: in.Name,
	})
	if err != nil {
		logs.CtxErrorf(ctx, "SayHello failed: %v", err)
		return nil, toKratosError(err)
	}
	return &v1.HelloReply{Message: "Hello " + resp.Hello}, nil
}

func toKratosError(err error) error {
	if se, ok := err.(errorx.StatusError); ok {
		return errors.New(int(se.Code()), "DOMAIN_ERROR", se.Msg())
	}
	return errors.InternalServer("INTERNAL_ERROR", err.Error())
}
