package greeter

import "kratos_template/internal/domain/greeter/entity"

type CreateGreeterRequest struct {
	Hello string
}

func (r *CreateGreeterRequest) ToEntity() *entity.Greeter {
	return &entity.Greeter{
		Hello: r.Hello,
	}
}

type GreeterResponse struct {
	ID    int64
	Hello string
}

func newResponse(g *entity.Greeter) *GreeterResponse {
	return &GreeterResponse{
		ID:    g.ID,
		Hello: g.Hello,
	}
}
