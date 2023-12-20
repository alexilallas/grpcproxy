package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/alexilallas/grpcproxy/internal/core/entity"
	"github.com/alexilallas/grpcproxy/internal/core/port"
	"github.com/alexilallas/grpcproxy/internal/core/port/mocks"
	"github.com/stretchr/testify/assert"
)

func TestProvideTaskUseCase(t *testing.T) {
	gateway := new(mocks.TaskGateway)

	t.Run("Should return an instance of port.TaskUseCase", func(t *testing.T) {
		instance := ProvideTaskUseCase(gateway)
		assert.NotNil(t, instance)
		assert.Implements(t, (*port.TaskUseCase)(nil), instance)
	})
}

func TestTaskUseCase_Create(t *testing.T) {
	var (
		ctx     = context.Background()
		gateway = new(mocks.TaskGateway)
		task    = ProvideTaskUseCase(gateway)
	)

	t.Run("Should return empty task when a generic error happen", func(t *testing.T) {
		gateway.On("CreateTask", ctx).Return(entity.Task{}, errors.New("generic error")).Once()

		actual, err := task.Create(ctx)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Task{}, actual)
	})

	t.Run("Should succeed", func(t *testing.T) {
		gateway.On("CreateTask", ctx).Return(entity.Task{}, nil)

		actual, err := task.Create(ctx)
		assert.Nil(t, err)
		assert.Equal(t, entity.Task{}, actual)
	})
}

func TestTaskUseCase_Get(t *testing.T) {
	var (
		ctx     = context.Background()
		id      = "task_id"
		gateway = new(mocks.TaskGateway)
		task    = ProvideTaskUseCase(gateway)
	)

	t.Run("Should return empty task when a generic error happen", func(t *testing.T) {
		gateway.On("GetTask", ctx, id).Return(entity.Task{}, errors.New("generic error")).Once()

		actual, err := task.Get(ctx, id)
		assert.NotNil(t, err)
		assert.Equal(t, entity.Task{}, actual)
	})

	t.Run("Should succeed", func(t *testing.T) {
		gateway.On("GetTask", ctx, id).Return(entity.Task{}, nil)

		actual, err := task.Get(ctx, id)
		assert.Nil(t, err)
		assert.Equal(t, entity.Task{}, actual)
	})
}
