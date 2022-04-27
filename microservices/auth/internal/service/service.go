package service

import (
	"github.com/Askalag/aska/microservices/auth/internal/provider"
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
)

type Service struct {
	Auth    Auth
	Session Session
}

type Session interface {
}

type Auth interface {
	Status() (string, error)
	SignIn(req *av1.SignInRequest) (*av1.SignInResponse, error)
	SignUp(req *repository.User) (*av1.SignUpResponse, error)
	FindUserByLogin(login string) (*repository.User, error)
	CreateUser(u *repository.User) (int, error)
}

func NewService(r *repository.Repo, p provider.Provider) *Service {
	return &Service{Auth: NewAuthService(&r.AuthRepo, &p)}
}
