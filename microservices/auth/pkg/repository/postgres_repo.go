package repository

import (
	"fmt"
	signIn_v1 "github.com/Askalag/protolib/gen/proto/go/sign_in/v1"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func (p *PostgresRepo) SignIn(req *signIn_v1.SignInRequest) (*signIn_v1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresRepo) Ping() error {
	return p.db.Ping()
}

func NewPostgresRepo(dbc DBConfig) *PostgresRepo {
	sql, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			dbc.Host, dbc.Port, dbc.Username, dbc.DBName, dbc.Password, dbc.SSLMode))
	if err != nil {
		log.Fatalln(err)
	}
	return &PostgresRepo{db: sql}
}
