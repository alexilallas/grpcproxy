package entity

import (
	"testing"

	"github.com/alexilallas/grpcproxy/internal/grpc"
	"github.com/stretchr/testify/assert"
)

func TestTask_ToOutput(t *testing.T) {

	t.Run("Should convert the entity to grpc type", func(t *testing.T) {
		task := Task{
			Id:         "id",
			Status:     "status",
			StatusCode: 200,
		}
		expected := grpc.Task{
			Id:             task.Id,
			Status:         task.Status,
			HttpStatusCode: task.StatusCode,
		}

		assert.Equal(t, &expected, task.ToOutput())
	})
}
