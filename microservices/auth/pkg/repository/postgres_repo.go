package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func (p *PostgresRepo) Ping() error {
	fmt.Println(*p.db)
	return p.db.Ping()
}

func NewPostgresRepo(dbc DBConfig) (*PostgresRepo, error) {
	sql, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			dbc.Host, dbc.Port, dbc.Username, dbc.DBName, dbc.Password, dbc.SSLMode))
	if err != nil {
		return nil, err
	}
	return &PostgresRepo{db: sql}, nil
}
