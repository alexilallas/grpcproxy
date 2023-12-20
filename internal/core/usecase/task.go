package usecase

import (
	"context"

	"github.com/alexilallas/grpcproxy/internal/core/entity"
	"github.com/alexilallas/grpcproxy/internal/core/port"
)

type TaskUseCase struct {
	gateway port.TaskGateway
}

func (t TaskUseCase) Create(ctx context.Context) (entity.Task, error) {
	return t.gateway.CreateTask(ctx)
}

func (t TaskUseCase) Get(ctx context.Context, id string) (entity.Task, error) {
	return t.gateway.GetTask(ctx, id)
}

func ProvideTaskUseCase(gateway port.TaskGateway) port.TaskUseCase {
	return TaskUseCase{gateway: gateway}
}
