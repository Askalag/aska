package service

import (
	"context"
	"github.com/askalag/aska/microservices/webapp/pkg"
	status_v1 "github.com/askalag/aska/microservices/webapp/proto/status/v1"
)

type History interface {
	GrpcStatus(ctx context.Context, req *status_v1.StatusRequest) (*status_v1.StatusResponse, error)
}

type Auth interface {
}

type Task interface {
}

type App interface {
	Status() map[string]interface{}
	StatusAll() map[string]interface{}
}

type Service struct {
	Auth    Auth
	History History
	Task    Task
	App     App
}

func NewService(tcp pkg.ServicesTCP) *Service {
	return &Service{
		App:     NewAppService(tcp),
		History: NewHistoryService(tcp),
	}
}
