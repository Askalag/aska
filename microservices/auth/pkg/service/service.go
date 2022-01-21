package service

import (
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	signIn_v1 "github.com/Askalag/protolib/gen/proto/go/sign_in/v1"
)

type Service struct {
	auth Auth
}

type Auth interface {
	SignIn(req *signIn_v1.SignInRequest) (*signIn_v1.SignInResponse, error)
}

func NewService(r *repository.Repo) *Service {
	return &Service{auth: NewAuthService(&r.AuthRepo)}
}
