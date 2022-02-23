package service

import (
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	siv1 "github.com/Askalag/protolib/gen/proto/go/signin/v1"
)

type AuthService struct {
	authRepo repository.AuthRepo
}

func (a *AuthService) Status() (string, error) {
	return "Auth service is alive", nil
}

func (a *AuthService) SignIn(req *siv1.SignInRequest) (*siv1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(r *repository.AuthRepo) *AuthService {
	return &AuthService{authRepo: *r}
}
