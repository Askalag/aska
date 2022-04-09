package server

import (
	"context"
	"github.com/Askalag/aska/microservices/auth/internal/service"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
	stv1 "github.com/Askalag/protolib/gen/proto/go/status/v1"
)

type AuthServer struct {
	auth service.Auth
}

func (s *AuthServer) SignUp(req *av1.SignUpRequest) (*av1.SignUpResponse, error) {
	_, err := s.auth.SignUp(req)
	if err != nil {
		return &av1.SignUpResponse{}, err
	}
	//TODO CONVERT
	//TODO implement me
	//TODO Start from here
	panic("implement me")
}

func (s *AuthServer) SignIn(req *av1.SignInRequest) (*av1.SignInResponse, error) {
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
