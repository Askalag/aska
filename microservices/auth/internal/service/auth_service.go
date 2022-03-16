package service

import (
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
)

type AuthService struct {
	authRepo repository.AuthRepo
}

func (a *AuthService) Status() (string, error) {
	return "Auth service is alive", nil
}

func (a *AuthService) SignIn(req *av1.SignInRequest) (*av1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) SignUp(req *av1.SignUpRequest) (*repository.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(r *repository.AuthRepo) *AuthService {
	return &AuthService{authRepo: *r}
}
