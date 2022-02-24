package repository

import (
	"fmt"
	si1 "github.com/Askalag/protolib/gen/proto/go/signin/v1"
)

type DBConfig struct {
	host     string
	port     string
	username string
	password string
	dbName   string
	sslMode  string
}

type Repo struct {
	AuthRepo AuthRepo
}

type UserRepo interface {
	getByUsername()
}

type AuthRepo interface {
	Ping() error
	SignIn(req *si1.SignInRequest) (*si1.SignInResponse, error)
}

func (c *DBConfig) BuildConnString(driver string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		driver, c.username, c.password, c.host, c.port, c.dbName, c.sslMode)
}

func NewDBConfig(h string, p string, u string, pass string, dbn string, ssl string) *DBConfig {
	return &DBConfig{
		host:     h,
		port:     p,
		username: u,
		password: pass,
		dbName:   dbn,
		sslMode:  ssl,
	}
}

func NewRepo(c *DBConfig) *Repo {
	return &Repo{AuthRepo: NewPostgresRepo(*c)}
}
