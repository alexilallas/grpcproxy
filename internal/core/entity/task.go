package entity

import "github.com/alexilallas/grpcproxy/internal/grpc"

type Task struct {
	Id         string `json:"id"`
	Status     string `json:"status"`
	StatusCode int32  `json:"httpStatusCode"`
}

func (t Task) ToOutput() *grpc.Task {
	return &grpc.Task{
		Id:             t.Id,
		Status:         t.Status,
		HttpStatusCode: t.StatusCode,
	}
}
