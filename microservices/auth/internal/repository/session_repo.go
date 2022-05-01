package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	sessionTable = "refresh_session"
	sessionAlive = time.Hour * 24 * 60 // ~ 60 days
)

type RefreshSession struct {
	Id           int       `db:"id"`
	UserId       int       `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	Ip           string    `db:"ip"`
	ExpiresIn    time.Time `db:"expires_in"`
	CreatedAt    time.Time `db:"created_at"`
}

type SessionRepository struct {
	db *sqlx.DB
}

// clearByUserId deleting all sessions by userId
func (s *SessionRepository) ClearByUserId(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s=$1", sessionTable, "user_id")
	_, err := s.db.Query(query, userId)
	return err
}

// CheckSession return false if a session is expired or not exists
func (s *SessionRepository) Check(refreshToken string) bool {
	session := &RefreshSession{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1 LIMIT 1", sessionTable, "refresh_token")
	err := s.db.Get(session, query, refreshToken)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		log.Errorf("error CheckSession with refreshToken: '%s'", refreshToken)
		return false
	}
	if time.Now().UTC().After(session.ExpiresIn) {
		return false
	}
	return true
}

func (s *SessionRepository) Create(userId int, ip string) (int, error) {
	session := &RefreshSession{
		UserId:       userId,
		RefreshToken: uuid.New().String(),
		ExpiresIn:    time.Now().Add(sessionAlive),
		Ip:           ip,
	}
	query := fmt.Sprintf(
		"INSERT INTO %s "+
			"(user_id, refresh_token, expires_in, ip) "+
			"VALUES ($1, $2, $3, $4) RETURNING id", sessionTable)
	err := s.db.QueryRow(query, session.UserId, session.RefreshToken, session.ExpiresIn.UTC(), session.Ip).Scan(&session.Id)
	return session.Id, err
}

func NewSessionRepo(db *sqlx.DB) *SessionRepository {
	return &SessionRepository{db: db}
}
