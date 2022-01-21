package service

import (
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	signIn_v1 "github.com/Askalag/protolib/gen/proto/go/sign_in/v1"
)

type AuthService struct {
	authRepo repository.AuthRepo
}

func NewAuthService(r *repository.AuthRepo) *AuthService {
	return &AuthService{authRepo: *r}
}

func (a *AuthService) SignIn(req *signIn_v1.SignInRequest) (*signIn_v1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}
