package repository

import (
	"database/sql"
	"fmt"
	av1 "github.com/Askalag/protolib/gen/proto/go/auth/v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	userTable = "users"

	errUserIsNotUnique = "the user is not unique"
	errCommonSql       = "unknown sql error"
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

type AuthRepository struct {
	db *sqlx.DB
}

func (p *AuthRepository) CreateUser(u *User) (int, error) {
	u.LastModified = time.Now()

	isUniq := p.isUserUnique(u)
	if !isUniq {
		return 0, status.Errorf(codes.InvalidArgument, errUserIsNotUnique)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s "+
			"(login, f_name, l_name, password, email, last_modified) "+
			"VALUES ($1, $2, $3, $4, $5, $6) "+
			"RETURNING id",
		userTable)

	var id int
	err := p.db.QueryRow(query, u.Login, u.FirstName, u.LastName, u.Password, u.Email, u.LastModified.UTC()).Scan(&id)
	if err != nil {
		return 0, status.Errorf(codes.InvalidArgument, errCommonSql)
	}
	return id, nil
}

func (p *AuthRepository) FindUserByEmail(email string) (*User, error) {
	u := &User{}

	query := fmt.Sprintf("select * from %s where %s=$1", userTable, "email")
	err := p.db.Get(u, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Unknown, err.Error())
	}
	return u, nil
}

func (p *AuthRepository) FindUserByLogin(login string) (*User, error) {
	u := &User{}

	query := fmt.Sprintf("select * from %s where %s=$1", userTable, "login")
	err := p.db.Get(u, query, login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Unknown, err.Error())
	}
	return u, nil
}

func (p *AuthRepository) SignIn(req *av1.SignInRequest) (*av1.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *AuthRepository) Ping() error {
	return p.db.Ping()
}

// isUserUnique return true if a new user has unique values. (see sql table)
// login, email - is unique
func (p *AuthRepository) isUserUnique(u *User) bool {
	f := &User{}
	query := fmt.Sprintf("SELECT * from %s where %s=$1 OR %s=$2",
		userTable, "login", "email")
	err := p.db.Get(f, query, u.Login, u.Email)
	if err != nil && err == sql.ErrNoRows {
		return true
	}
	return false
}

func NewAuthRepo(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
