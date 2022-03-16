package repository

import (
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
	id           int       `db:"id"`
	login        string    `db:"login"`
	firstName    string    `db:"f_name"`
	lastName     string    `db:"l_name"`
	password     string    `db:"password"`
	email        string    `db:"email"`
	active       bool      `db:"active"`
	dateCreated  time.Time `db:"date_created"`
	lastModified time.Time `db:"last_modified"`
}

type PostgresRepo struct {
	db *sqlx.DB
}

func (p *PostgresRepo) SignUp(req *av1.SignUpRequest) (*User, error) {
	var query = fmt.Sprintf(
		`insert into %v (login, firstName, lastName, password, email) 
					values (%v, %v, %v, %v, %v)`,
		userTable,
		req.Login,
		req.FirstName,
		req.LastName,
		req.Password,
		req.Email,
	)

	user := &User{}
	u, err := p.db.Query(query)
	if err != nil {
		return user, err
	}

	err = u.Scan(user)
	if err != nil {
		return user, err
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
