package repository

import (
	"database/sql"
	"fmt"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	userTable = "users"
)

type User struct {
	Id           int       `db:"id"`
	Login        string    `db:"login"`
	FirstName    string    `db:"f_name"`
	LastName     string    `db:"l_name"`
	Password     string    `db:"password"`
	Email        string    `db:"email"`
	Active       bool      `db:"active"`
	DateCreated  time.Time `db:"date_created"`
	LastModified time.Time `db:"last_modified"`
}

type PostgresRepo struct {
	db *sqlx.DB
}

func (p *PostgresRepo) CreateUser(u *User) (int, error) {
	u.LastModified = time.Now()
	query := fmt.Sprintf(
		"INSERT INTO %s "+
			"(login, f_name, l_name, password, email, last_modified) "+
			"VALUES "+
			"('%v', '%v', '%v', '%v', '%v', '%v') "+
			"RETURNING id",
		userTable, u.Login, u.FirstName, u.LastName, u.Password, u.Email, u.LastModified.UTC().Format(time.RFC3339Nano))
	var id int // todo
	err := p.db.QueryRow(query).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostgresRepo) FindUserByEmail(email string) (*User, error) {
	u := &User{}
	err := p.db.Get(u, "select * from users where email=$1", email)
	return u, err
}

func (p *PostgresRepo) FindUserByLogin(login string) (*User, error) {
	u := &User{}
	err := p.db.Get(u, "select * from users where login=$1 LIMIT 1", login)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (p *PostgresRepo) SignIn(req *av1.SignInRequest) (*av1.SignInResponse, error) {
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
