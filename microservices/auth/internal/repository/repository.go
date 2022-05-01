package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
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
	AuthRepo    AuthRepo
	SessionRepo SessionRepo
}

type SessionRepo interface {
	Create(userId int, ip string) (int, error)
	Check(uuid string) bool
	ClearByUserId(userId int) error
}

type AuthRepo interface {
	Ping() error
	FindUserByLogin(login string) (*User, error)
	FindUserByEmail(email string) (*User, error)
	CreateUser(u *User) (int, error)
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
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			c.host, c.port, c.username, c.dbName, c.password, c.sslMode))
	if err != nil {
		log.Fatalln(err)
	}
	return &Repo{
		AuthRepo:    NewAuthRepo(db),
		SessionRepo: NewSessionRepo(db),
	}
}
