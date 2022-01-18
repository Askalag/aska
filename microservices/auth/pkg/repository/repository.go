package repository

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Repo struct {
	AuthRepo AuthRepo
}

type AuthRepo interface {
	Ping() error
}

func NewRepo(c *DBConfig) (*Repo, error) {
	authDb, err := NewPostgresRepo(*c)
	if err != nil {
		return nil, err
	}

	return &Repo{AuthRepo: authDb}, nil
}
