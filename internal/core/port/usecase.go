package port

import (
	"context"

	"github.com/alexilallas/grpcproxy/internal/core/entity"
)

type TaskUseCase interface {
	Create(ctx context.Context) (entity.Task, error)
	Get(ctx context.Context, id string) (entity.Task, error)
}
