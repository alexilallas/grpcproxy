package inbound

import (
	"context"
	"errors"
	"testing"

	"github.com/alexilallas/grpcproxy/internal/core/entity"
	"github.com/alexilallas/grpcproxy/internal/core/port/mocks"
	"github.com/alexilallas/grpcproxy/internal/grpc"
	"github.com/stretchr/testify/assert"
)

func TestProvideNewHandler(t *testing.T) {
	usecase := new(mocks.TaskUseCase)

	t.Run("Should return an instance of grpc.TaskManagerServer", func(t *testing.T) {
		instance := ProvideNewHandler(usecase)
		assert.NotNil(t, instance)
		assert.Implements(t, (*grpc.TaskManagerServer)(nil), instance)
	})
}

func TestTaskHandler_CreateTask(t *testing.T) {
	var (
		ctx         = context.Background()
		taskUseCase = new(mocks.TaskUseCase)
		handler     = ProvideNewHandler(taskUseCase)
	)

	t.Run("Should return an error when usecase return an error", func(t *testing.T) {
		taskUseCase.On("Create", ctx).Return(entity.Task{}, errors.New("generic error")).Once()
		actual, err := handler.CreateTask(ctx, nil)
		assert.NotNil(t, err)
		assert.IsType(t, &grpc.Task{}, actual)
	})

	t.Run("Should succeed", func(t *testing.T) {
		taskUseCase.On("Create", ctx).Return(entity.Task{}, nil)
		actual, err := handler.CreateTask(ctx, nil)
		assert.Nil(t, err)
		assert.IsType(t, &grpc.Task{}, actual)
	})
}

func TestTaskHandler_GetTask(t *testing.T) {
	var (
		ctx         = context.Background()
		taskUseCase = new(mocks.TaskUseCase)
		handler     = ProvideNewHandler(taskUseCase)
		task        = grpc.Task{Id: "id"}
	)

	t.Run("Should return an error when usecase return an error", func(t *testing.T) {
		taskUseCase.On("Get", ctx, task.Id).Return(entity.Task{}, errors.New("generic error")).Once()
		actual, err := handler.GetTask(ctx, &task)
		assert.NotNil(t, err)
		assert.IsType(t, &grpc.Task{}, actual)
	})

	t.Run("Should succeed", func(t *testing.T) {
		taskUseCase.On("Get", ctx, task.Id).Return(entity.Task{}, nil)
		actual, err := handler.GetTask(ctx, &task)
		assert.Nil(t, err)
		assert.IsType(t, &grpc.Task{}, actual)
	})
}
