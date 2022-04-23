package provider

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var sha = "HS256"

type Provider interface {
	CreateToken(u *repository.User) (string, error)
	ParseAndVerifyToken(tokenString string) (*jwt.Token, error)
	CreateRefreshToken() string
}

type JWTProvider struct {
	Provider
	config JWTConfig
	repo   repository.AuthRepo
}

// JWTConfig
type JWTConfig struct {
	secret     string
	signKey    *rsa.PrivateKey
	alg        string
	tokenAlive int // tokenAlive - in minutes
}

type AuthClaims struct {
	*jwt.StandardClaims
	TokenType string
	UserInfo  *repository.User
}

func (p *JWTProvider) VerifyPasswordHash(pass, passHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(pass))
	return err == nil
}

func (p *JWTProvider) HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	return string(bytes), err
}

// ParseAndVerifyToken param tokenString should be without "Bearer " ...
func (p *JWTProvider) ParseAndVerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong signed token method")
		}
		return []byte(p.config.secret), nil
	})
	return token, err
}

func (p *JWTProvider) CreateToken(u *repository.User) (string, error) {
	if u.Login == "" {
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

func (p *JWTProvider) CreateRefreshToken() string {
	arr, err := generateRndBytesArr(32)
	if err != nil {
		panic("err creating a new refresh token")
	}
	return string(arr)
}

func buildClaims(u *repository.User, c JWTConfig) *AuthClaims {
	return &AuthClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(c.tokenAlive)).Unix(),
		},
		"",
		u,
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

func NewJWTProvider(c *JWTConfig, repo *repository.AuthRepo) *JWTProvider {
	return &JWTProvider{
		config: *c,
		repo:   *repo,
	}
}
