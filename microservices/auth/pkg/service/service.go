package service

import (
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	auth_v1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
)

type Service struct {
	auth Auth
}

type Auth interface {
	SignIn(req *auth_v1.SignInRequest) (*auth_v1.TokenResponse, error)
}

func NewService(r *repository.Repo) *Service {
	return &Service{auth: NewAuthService(&r.AuthRepo)}
}
