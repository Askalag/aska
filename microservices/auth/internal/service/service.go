package service

import (
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	signinv1 "github.com/Askalag/protolib/gen/proto/go/signin/v1"
)

type Service struct {
	Auth Auth
}

type Auth interface {
	SignIn(req *signinv1.SignInRequest) (*signinv1.SignInResponse, error)
	Status() (string, error)
}

func NewService(r *repository.Repo) *Service {
	return &Service{Auth: NewAuthService(&r.AuthRepo)}
}
