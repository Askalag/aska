package provider

import "github.com/Askalag/aska/microservices/auth/pkg/repository"

type JWTProvider interface {
}

type Provider struct {
	config JWTConfig
	repo   repository.AuthRepo
}

type JWTConfig struct {
	secret string
}

func BuildJWTConfig(s string) *JWTConfig {
	return &JWTConfig{secret: s}
}

func NewJWTProvider(c *JWTConfig, repo *repository.AuthRepo) *Provider {
	return &Provider{
		config: *c,
		repo:   *repo,
	}
}
