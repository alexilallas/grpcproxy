package inbound

import (
	"context"

	"github.com/alexilallas/grpcproxy/internal/core/port"
	"github.com/alexilallas/grpcproxy/internal/grpc"
	"github.com/golang/protobuf/ptypes/empty"
)

type TaskHandler struct {
	grpc.UnimplementedTaskManagerServer
	task port.TaskUseCase
}

func (h TaskHandler) CreateTask(ctx context.Context, _ *empty.Empty) (*grpc.Task, error) {
	task, err := h.task.Create(ctx)
	if err != nil {
		return nil, err
	}
	return task.ToOutput(), nil
}

func (h TaskHandler) GetTask(ctx context.Context, task *grpc.Task) (*grpc.Task, error) {
	response, err := h.task.Get(ctx, task.GetId())
	if err != nil {
		return nil, err
	}
	return response.ToOutput(), nil
}

func ProvideNewHandler(usecase port.TaskUseCase) grpc.TaskManagerServer {
	return TaskHandler{task: usecase}
}
