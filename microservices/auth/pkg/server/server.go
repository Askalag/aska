package server

import (
	"context"
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	"github.com/Askalag/aska/microservices/auth/pkg/service"
	auth_v1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
)

type Server struct {
	auth Auth
}

type Auth interface {
	SignIn(req *auth_v1.SignInRequest) (*auth_v1.TokenResponse, error)
}

func (s *Server) SignIn(context.Context, *auth_v1.SignInRequest) (*auth_v1.TokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewServer(r *repository.Repo) (*Server, error) {
	return &Server{auth: service.NewAuthService(&r.AuthRepo)}, nil
}
