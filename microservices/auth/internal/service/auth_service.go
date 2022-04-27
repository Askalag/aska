package service

import (
	"fmt"
	"github.com/Askalag/aska/microservices/auth/internal/provider"
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
)

type AuthService struct {
	authRepo     repository.AuthRepo
	authProvider provider.Provider
}

func (a *AuthService) FindUserByLogin(login string) (*repository.User, error) {
	return a.authRepo.FindUserByLogin(login)
}

func (a *AuthService) CreateUser(u *repository.User) (int, error) {
	pswHash, err := a.authProvider.HashPassword(u.Password)
	if err != nil {
		return 0, err
	}
	u.Password = pswHash
	return a.authRepo.CreateUser(u)
}

func (a *AuthService) SignIn(u *av1.SignInRequest) (*av1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) SignUp(u *repository.User) (*av1.SignUpResponse, error) {
	f, err := a.FindUserByLogin(u.Login)
	if err != nil {
		return nil, err
	}
	if f != nil && u.Login == f.Login {
		return nil, fmt.Errorf("The User with login: '%s' is already exists", u.Login)
	}

	id, err := a.CreateUser(u)
	if err != nil {
		return nil, err
	}
	u.Id = id

	token, err := a.authProvider.CreateToken(u)
	if err != nil {
		return nil, err
	}

	return &av1.SignUpResponse{
		Token:        token,
		RefreshToken: "",
	}, nil
}

func (a *AuthService) Status() (string, error) {
	return "Auth service is alive", nil
}

func NewAuthService(r *repository.AuthRepo, p *provider.Provider) *AuthService {
	return &AuthService{
		authRepo:     *r,
		authProvider: *p,
	}
}
