package port

import (
	"context"

	"github.com/alexilallas/grpcproxy/internal/core/entity"
)

type TaskGateway interface {
	CreateTask(ctx context.Context) (entity.Task, error)
	GetTask(ctx context.Context, id string) (entity.Task, error)
}
