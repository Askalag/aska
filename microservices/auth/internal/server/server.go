package server

import (
	"context"
	"github.com/Askalag/aska/microservices/auth/internal/service"
	siv1 "github.com/Askalag/protolib/gen/proto/go/signin/v1"
	stv1 "github.com/Askalag/protolib/gen/proto/go/status/v1"
)

type Server struct {
	auth Auth
}

type Auth interface {
	SignIn(req *siv1.SignInRequest) (*siv1.SignInResponse, error)
	Status(ctx context.Context, res *stv1.StatusRequest) (*stv1.StatusResponse, error)
}

func NewServer(s *service.Service) (*Server, error) {
	return &Server{auth: NewAuthServer(&s.Auth)}, nil
}
