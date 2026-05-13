package service

import (
	"context"

	"github.com/Martindeeepdark/go-common/errorx"
	"github.com/Martindeeepdark/go-common/logs"

	"kratos_template/internal/domain/greeter/entity"
	"kratos_template/internal/domain/greeter/repository"
)

var (
	ErrGreeterNotFound = errorx.New(10001)
)

func init() {
	errorx.Register(10001, "greeter not found")
}

type GreeterService struct {
	repo repository.GreeterRepository
}

func NewGreeterService(repo repository.GreeterRepository) *GreeterService {
	return &GreeterService{repo: repo}
}

func (s *GreeterService) Create(ctx context.Context, g *entity.Greeter) (*entity.Greeter, error) {
	logs.CtxInfof(ctx, "Creating greeter: %s", g.Hello)
	return s.repo.Save(ctx, g)
}
