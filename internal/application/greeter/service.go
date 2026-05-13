package greeter

import (
	"context"

	"kratos_template/internal/domain/greeter/service"
)

type AppService struct {
	greeterSvc *service.GreeterService
}

func NewAppService(greeterSvc *service.GreeterService) *AppService {
	return &AppService{greeterSvc: greeterSvc}
}

func (s *AppService) CreateGreeter(ctx context.Context, req *CreateGreeterRequest) (*GreeterResponse, error) {
	result, err := s.greeterSvc.Create(ctx, req.ToEntity())
	if err != nil {
		return nil, err
	}
	return newResponse(result), nil
}

func (s *AppService) GetGreeter(ctx context.Context, id int64) (*GreeterResponse, error) {
	return nil, nil
}
