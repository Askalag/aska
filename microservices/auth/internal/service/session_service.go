package service

import (
	"github.com/Askalag/aska/microservices/auth/internal/repository"
)

type SessionService struct {
	sessionRepo repository.SessionRepo
}

func (s *SessionService) create(userId int) (*repository.RefreshSession, error) {
	return nil, nil
}
