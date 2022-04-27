package service

import (
	"github.com/Askalag/aska/microservices/auth/internal/repository"
)

type SessionService struct {
	sessionRepo repository.SessionRepo
}
