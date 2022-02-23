package repository

import (
	"fmt"
	siv1 "github.com/Askalag/protolib/gen/proto/go/signin/v1"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func (p *PostgresRepo) SignIn(req *siv1.SignInRequest) (*siv1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresRepo) Ping() error {
	return p.db.Ping()
}

func NewPostgresRepo(dbc DBConfig) *PostgresRepo {
	sql, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			dbc.host, dbc.port, dbc.username, dbc.dbName, dbc.password, dbc.sslMode))
	if err != nil {
		log.Fatalln(err)
	}
	return &PostgresRepo{db: sql}
}
