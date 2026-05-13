package repository

import (
	"context"

	"kratos_template/internal/domain/greeter/entity"
)

type GreeterRepository interface {
	Save(ctx context.Context, g *entity.Greeter) (*entity.Greeter, error)
	Update(ctx context.Context, g *entity.Greeter) (*entity.Greeter, error)
	FindByID(ctx context.Context, id int64) (*entity.Greeter, error)
	ListAll(ctx context.Context) ([]*entity.Greeter, error)
}
