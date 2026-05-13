package infrastructure

import (
	"context"

	"gorm.io/gorm"

	"kratos_template/infrastructure/model"
	"kratos_template/internal/domain/greeter/entity"
	"kratos_template/internal/domain/greeter/repository"
)

type greeterRepo struct {
	data *Data
}

func NewGreeterRepo(data *Data) repository.GreeterRepository {
	return &greeterRepo{data: data}
}

func toGreeterModel(e *entity.Greeter) *model.Greeter {
	return &model.Greeter{
		ID:    e.ID,
		Hello: e.Hello,
	}
}

func toGreeterEntity(m *model.Greeter) *entity.Greeter {
	return &entity.Greeter{
		ID:    m.ID,
		Hello: m.Hello,
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *entity.Greeter) (*entity.Greeter, error) {
	m := toGreeterModel(g)
	if err := r.data.DB().WithContext(ctx).Create(m).Error; err != nil {
		return nil, err
	}
	g.ID = m.ID
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *entity.Greeter) (*entity.Greeter, error) {
	m := toGreeterModel(g)
	if err := r.data.DB().WithContext(ctx).Save(m).Error; err != nil {
		return nil, err
	}
	return g, nil
}

func (r *greeterRepo) FindByID(ctx context.Context, id int64) (*entity.Greeter, error) {
	var m model.Greeter
	if err := r.data.DB().WithContext(ctx).First(&m, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return toGreeterEntity(&m), nil
}

func (r *greeterRepo) ListAll(ctx context.Context) ([]*entity.Greeter, error) {
	var models []*model.Greeter
	if err := r.data.DB().WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	result := make([]*entity.Greeter, 0, len(models))
	for _, m := range models {
		result = append(result, toGreeterEntity(m))
	}
	return result, nil
}
