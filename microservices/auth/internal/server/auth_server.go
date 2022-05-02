package server

import (
	"context"
	conv "github.com/Askalag/aska/microservices/auth/internal/convertor"
	"github.com/Askalag/aska/microservices/auth/internal/service"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
	stv1 "github.com/Askalag/protolib/gen/proto/go/status/v1"
)

type AuthServer struct {
	auth service.Auth
}

func (s *AuthServer) RefreshToken(ctx context.Context, req *av1.RefreshTokenRequest) (*av1.RefreshTokenResponse, error) {
	return s.auth.RefreshTokenPair(req)
}

func (s *AuthServer) SignUp(ctx context.Context, req *av1.SignUpRequest) (*av1.SignUpResponse, error) {
	user := conv.SignUpRequestToUserV1(req)
	return s.auth.SignUp(user)
}

func (s *AuthServer) SignIn(ctx context.Context, req *av1.SignInRequest) (*av1.SignInResponse, error) {
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
