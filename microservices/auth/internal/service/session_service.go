package service

import (
	"github.com/Askalag/aska/microservices/auth/internal/repository"
)

type SessionService struct {
	sessionRepo repository.SessionRepo
}

func (s *SessionService) Create(userId int, ip string) (int, error) {
	return s.sessionRepo.Create(userId, ip)
}

func (s *SessionService) Check(uuid string) bool {
	return s.sessionRepo.Check(uuid)
}

func (s *SessionService) ClearByUserId(userId int) error {
	return s.sessionRepo.ClearByUserId(userId)
}

func NewSessionService(r *repository.SessionRepo) *SessionService {
	return &SessionService{
		sessionRepo: *r,
	}
}
