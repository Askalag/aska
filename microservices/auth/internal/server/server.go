package server

import (
	"context"
	"github.com/Askalag/aska/microservices/auth/internal/service"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
	stv1 "github.com/Askalag/protolib/gen/proto/go/status/v1"
)

type Server struct {
	Auth Auth
}

type Auth interface {
	Status(ctx context.Context, res *stv1.StatusRequest) (*stv1.StatusResponse, error)
	SignUp(req *av1.SignUpRequest) (*av1.SignUpResponse, error)
	SignIn(req *av1.SignInRequest) (*av1.SignInResponse, error)
}

func NewServer(s *service.Service) *Server {
	return &Server{Auth: NewAuthServer(&s.Auth)}
}
