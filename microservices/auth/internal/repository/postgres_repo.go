package repository

import (
	"errors"
	"fmt"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	userTable = "user"
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

func (p *PostgresRepo) FindUserByEmail(email string) (*User, error) {
	u := &User{}
	err := p.db.Get(u, "select * from aska_db.public.users where email=$1", email)
	return u, err
}

func (p *PostgresRepo) FindUserByLogin(login string) (*User, error) {
	var u *User
	err := p.db.Get(&u, "select * from aska_db.public.users where login=$1", login)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, fmt.Errorf("The User with login: '%s' not found", login)
	}
	return u, nil
}

func (p *PostgresRepo) SignUp(u *User) (*User, error) {
	user, err := p.FindUserByLogin(u.Login)
	if err != nil {
		return nil, err
	}
	if user.Login == u.Login {
		return nil, fmt.Errorf("The User with login: '%s' is already exists", u.Login)
	}

	var query = fmt.Sprintf(
		`insert into %v (login, firstName, lastName, password, email) 
					values (%v, %v, %v, %v, %v)`,
		userTable,
		u.Login,
		u.FirstName,
		u.LastName,
		u.Password,
		u.Email,
	)

	result, err := p.db.Query(query)
	if err != nil {
		return user, err
	}

	err = result.Scan(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
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
