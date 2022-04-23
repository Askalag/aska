package service

import (
	"github.com/Askalag/aska/microservices/auth/internal/provider"
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
)

type AuthService struct {
	authRepo     repository.AuthRepo
	authProvider provider.Provider
}

func (a *AuthService) Status() (string, error) {
	return "Auth service is alive", nil
}

func (a *AuthService) SignIn(u *av1.SignInRequest) (*av1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) SignUp(u *repository.User) (*av1.SignUpResponse, error) {
	u, err := a.authRepo.SignUp(u)
	if err != nil {
		return nil, err
	}

	token, err := a.authProvider.CreateToken(u)
	if err != nil {
		return nil, err
	}

	return &av1.SignUpResponse{
		Token:        token,
		RefreshToken: "",
	}, nil
}

func NewAuthService(r *repository.AuthRepo, p *provider.Provider) *AuthService {
	return &AuthService{
		authRepo:     *r,
		authProvider: *p,
	}
}
