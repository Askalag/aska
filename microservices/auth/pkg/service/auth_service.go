package service

import (
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	auth_v1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
)

type AuthService struct {
	authRepo repository.AuthRepo
}

func NewAuthService(r *repository.AuthRepo) *AuthService {
	return &AuthService{authRepo: *r}
}

func (a *AuthService) SignIn(req *auth_v1.SignInRequest) (*auth_v1.TokenResponse, error) {
	//TODO implement me
	panic("implement me")
}
