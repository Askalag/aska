package repository

import signIn_v1 "github.com/Askalag/protolib/gen/proto/go/sign_in/v1"

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
	SignIn(req *signIn_v1.SignInRequest) (*signIn_v1.SignInResponse, error)
}

func NewRepo(c *DBConfig) *Repo {
	return &Repo{AuthRepo: NewPostgresRepo(*c)}
}
