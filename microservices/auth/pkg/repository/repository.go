package repository

import auth_v1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Repo struct {
	AuthRepo AuthRepo
}

type AuthRepo interface {
	Ping() error
	SignIn(req *auth_v1.SignInRequest) (*auth_v1.TokenResponse, error)
}

func NewRepo(c *DBConfig) *Repo {
	return &Repo{AuthRepo: NewPostgresRepo(*c)}
}
