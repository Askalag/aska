package server

import (
	"context"
	"github.com/Askalag/aska/microservices/auth/pkg/service"
	av1 "github.com/Askalag/aska/microservices/auth/proto/auth/v1"
	stv1 "github.com/Askalag/aska/microservices/auth/proto/status/v1"
)

type Server struct {
	Auth Auth
}

type Auth interface {
	Status(ctx context.Context, res *stv1.StatusRequest) (*stv1.StatusResponse, error)
	SignUp(ctx context.Context, req *av1.SignUpRequest) (*av1.SignUpResponse, error)
	SignIn(ctx context.Context, req *av1.SignInRequest) (*av1.SignInResponse, error)
	RefreshToken(ctx context.Context, req *av1.RefreshTokenRequest) (*av1.RefreshTokenResponse, error)
}

func NewServer(s *service.Service) *Server {
	return &Server{Auth: NewAuthServer(&s.Auth)}
}
