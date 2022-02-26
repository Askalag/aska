package provider

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	"github.com/golang-jwt/jwt"
	"time"
)

var sha = "HS256"

type JWTProvider interface {
}

type Provider struct {
	config JWTConfig
	repo   repository.AuthRepo
}

// JWTConfig tokenAlive - minutes
type JWTConfig struct {
	secret     string
	signKey    *rsa.PrivateKey
	alg        string
	tokenAlive int
}

type UserInfo struct {
	Username string
}

type AuthClaims struct {
	*jwt.StandardClaims
	TokenType string
	UserInfo  UserInfo
}

// ParseAndVerifyToken param tokenString should be without "Bearer " ...
func (p *Provider) ParseAndVerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong signed token method")
		}
		return []byte(p.config.secret), nil
	})
	return token, err
}

func (p *Provider) CreateToken(u *UserInfo) (string, error) {
	if u.Username == "" {
		return "", errors.New("err creating a new token, username invalid")
	}

	token := jwt.New(jwt.GetSigningMethod(p.config.alg))
	token.Claims = buildClaims(u, p.config)

	signedString, err := token.SignedString(p.config.signKey)
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func (p *Provider) CreateRefreshToken() string {
	arr, err := generateRndBytesArr(32)
	if err != nil {
		panic("err creating a new refresh token")
	}
	return string(arr)
}

func buildClaims(u *UserInfo, c JWTConfig) *AuthClaims {
	return &AuthClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(c.tokenAlive)).Unix(),
		},
		"",
		*u,
	}
}

func generateRndBytesArr(count int) ([]byte, error) {
	b := make([]byte, count)
	_, err := rand.Read(b)
	if err != nil {
		return b, err
	}
	return b, nil
}

func generateKey(secret string) *rsa.PrivateKey {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(secret))
	if err != nil {
		panic("generateKey")
	}
	return key
}

func BuildJWTConfig(secret string, tokenAlive int) *JWTConfig {
	return &JWTConfig{
		secret:     secret,
		signKey:    generateKey(secret),
		alg:        sha,
		tokenAlive: tokenAlive,
	}
}

func NewJWTProvider(c *JWTConfig, repo *repository.AuthRepo) *Provider {
	return &Provider{
		config: *c,
		repo:   *repo,
	}
}
