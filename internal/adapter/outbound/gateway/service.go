package gateway

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/alexilallas/grpcproxy/internal/core/entity"
	"github.com/alexilallas/grpcproxy/internal/core/port"
)

const (
	taskEndpoint         = "/task"
	headerAuthentication = "Authentication"
	tokenType            = "Basic"
)

type Service struct {
	client http.Client
	url    string
	token  string
}

func (s Service) CreateTask(ctx context.Context) (entity.Task, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.url+taskEndpoint, nil)
	if err != nil {
		return entity.Task{}, err
	}

	req.Header.Set(headerAuthentication, tokenType+" "+s.token)
	res, err := s.client.Do(req)
	if err != nil {
		return entity.Task{}, err
	}

	if res.StatusCode != http.StatusOK {
		return entity.Task{}, ErrorResponse(res.StatusCode)
	}
	var task entity.Task
	if err = json.NewDecoder(res.Body).Decode(&task); err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (s Service) GetTask(ctx context.Context, id string) (entity.Task, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.url+taskEndpoint+"/"+id, nil)
	if err != nil {
		return entity.Task{}, err
	}

	req.Header.Set(headerAuthentication, tokenType+" "+s.token)
	res, err := s.client.Do(req)
	if err != nil {
		return entity.Task{}, err
	}

	if res.StatusCode != http.StatusOK {
		return entity.Task{}, ErrorResponse(res.StatusCode)
	}
	var task entity.Task
	if err = json.NewDecoder(res.Body).Decode(&task); err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func ProvideTaskGateway(client http.Client, url, token string) port.TaskGateway {
	return Service{client: client, url: url, token: token}
}
