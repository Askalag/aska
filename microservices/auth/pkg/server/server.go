package server

import (
	"context"
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	"github.com/Askalag/aska/microservices/auth/pkg/service"
	signIn_v1 "github.com/Askalag/protolib/gen/proto/go/sign_in/v1"
)

type Server struct {
	auth Auth
}

type Auth interface {
	SignIn(req *signIn_v1.SignInRequest) (*signIn_v1.SignInResponse, error)
}

func (s *Server) SignIn(context.Context, *signIn_v1.SignInRequest) (*signIn_v1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewServer(r *repository.Repo) (*Server, error) {
	return &Server{auth: service.NewAuthService(&r.AuthRepo)}, nil
}
