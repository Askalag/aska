package service

import (
	"github.com/Askalag/aska/microservices/auth/internal/provider"
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	errCreateToken               = "error create token"
	errCommonSessionRepo         = "error common session repo"
	errRefTokenExpiredOrNotFound = "error refresh token is expired or not found"
)

type AuthService struct {
	authRepo       repository.AuthRepo
	sessionService Session
	authProvider   provider.Provider
}

func (a *AuthService) RefreshTokenPair(req *av1.RefreshTokenRequest) (*av1.RefreshTokenResponse, error) {
	session, err := a.sessionService.GetSessionByRefToken(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errCommonSessionRepo)
	}
	if session == nil || time.Now().UTC().After(session.ExpiresIn) {
		return nil, status.Errorf(codes.NotFound, errRefTokenExpiredOrNotFound)
	}

	u, err := a.authRepo.FindUserById(session.UserId)
	if err != nil {
		return nil, err
	}

	token, err := a.authProvider.CreateToken(u)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errCreateToken)
	}

	refToken, err := a.sessionService.DeleteByIdAndCreate(session.Id, u.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errCommonSessionRepo)
	}

	return &av1.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refToken.RefreshToken,
	}, nil
}

func (a *AuthService) FindUserByLogin(login string) (*repository.User, error) {
	return a.authRepo.FindUserByLogin(login)
}

func (a *AuthService) CreateUser(u *repository.User) (int, error) {
	pswHash, err := a.authProvider.HashPassword(u.Password)
	if err != nil {
		return 0, status.Errorf(codes.InvalidArgument, err.Error())
	}
	u.Password = pswHash
	return a.authRepo.CreateUser(u)
}

func (a *AuthService) SignIn(u *av1.SignInRequest) (*av1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) SignUp(u *repository.User) (*av1.SignUpResponse, error) {
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

func NewAuthService(r repository.AuthRepo, p provider.Provider, s Session) *AuthService {
	return &AuthService{
		authRepo:       r,
		sessionService: s,
		authProvider:   p,
	}
}
