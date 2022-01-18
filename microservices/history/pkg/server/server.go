package server

import (
	"context"
	"github.com/Askalag/aska/microservices/history/pkg/service"
	status_v1 "github.com/Askalag/protolib/gen/proto/go/status/v1"
)

type Server struct {
	Hs History
}

type History interface {
	Status(_ context.Context, _ *status_v1.StatusRequest) (*status_v1.StatusResponse, error)
}

func NewServer(service service.Service) *Server {
	return &Server{Hs: NewHistoryServer(service.HSService)}
}
