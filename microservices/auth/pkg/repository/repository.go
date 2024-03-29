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
	schema   string
}

type Repo struct {
	AuthRepo    AuthRepo
	SessionRepo SessionRepo
}

type SessionRepo interface {
	Create(userId int, ip string) (*RefreshSession, error)
	GetSessionByRefToken(refreshToken string) (*RefreshSession, error)
	ClearByUserId(userId int) error
	GetById(sessionId int) (*RefreshSession, error)
	DeleteById(sessionId int) error
}

type AuthRepo interface {
	Ping() error
	FindUserById(id int) (*User, error)
	FindUserByLogin(login string) (*User, error)
	FindUserByEmail(email string) (*User, error)
	CreateUser(u *User) (int, error)
}

func (c *DBConfig) BuildConnString(driver string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s&search_path=%s",
		driver, c.username, c.password, c.host, c.port, c.dbName, c.sslMode, c.schema)
}

func NewDBConfig(h string, p string, u string, pass string, dbn string, ssl string, schema string) *DBConfig {
	return &DBConfig{
		host:     h,
		port:     p,
		username: u,
		password: pass,
		dbName:   dbn,
		sslMode:  ssl,
		schema:   schema,
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
