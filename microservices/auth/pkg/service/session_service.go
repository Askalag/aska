package service

import (
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SessionService struct {
	sessionRepo repository.SessionRepo
}

func (s *SessionService) DeleteByIdAndCreate(oldSessionId int, userId int) (*repository.RefreshSession, error) {
	err := s.sessionRepo.DeleteById(oldSessionId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errCommonSessionRepo)
	}

	rs, err := s.sessionRepo.Create(userId, "0.0.0.0")
	if err != nil || rs == nil {
		return nil, status.Errorf(codes.Internal, errCommonSessionRepo)
	}
	return rs, nil
}

func (s *SessionService) Create(userId int, ip string) (*repository.RefreshSession, error) {
	return s.sessionRepo.Create(userId, ip)
}

func (s *SessionService) GetSessionByRefToken(refreshToken string) (*repository.RefreshSession, error) {
	return s.sessionRepo.GetSessionByRefToken(refreshToken)
}

func (s *SessionService) ClearByUserId(userId int) error {
	return s.sessionRepo.ClearByUserId(userId)
}

func NewSessionService(r *repository.SessionRepo) *SessionService {
	return &SessionService{
		sessionRepo: *r,
	}
}
