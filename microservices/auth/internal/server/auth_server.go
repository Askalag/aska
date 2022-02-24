package server

import (
	"context"
	"github.com/Askalag/aska/microservices/auth/internal/service"
	siv1 "github.com/Askalag/protolib/gen/proto/go/signin/v1"
	stv1 "github.com/Askalag/protolib/gen/proto/go/status/v1"
)

type AuthServer struct {
	auth service.Auth
}

func (s *AuthServer) SignIn(req *siv1.SignInRequest) (*siv1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *AuthServer) Status(_ context.Context, _ *stv1.StatusRequest) (*stv1.StatusResponse, error) {
	st, err := s.auth.Status()
	if err != nil {
		return nil, err
	}
	status := &stv1.StatusResponse{Status: st}
	return status, nil
}

func NewAuthServer(s *service.Auth) *AuthServer {
	return &AuthServer{auth: *s}
}
