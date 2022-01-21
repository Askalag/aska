package main

import (
	"flag"
	"fmt"
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

type AuthConfig struct {
	AuthAddr string
	AuthPort string
	DBConfig repository.DBConfig
}

const LogFile = "/tmp/auth_log.log"

func main() {
	c := buildConfig()

	initLog()

	initMigration(c.DBConfig)

	_ = repository.NewRepo(&c.DBConfig)

}

func buildConfig() *AuthConfig {
	authAddr := flag.String("auth_a", "", "http sign_in server address")
	authPort := flag.String("auth_p", "", "http sign_in port address")

	dbHost := flag.String("dbh", "", "db host address")
	dbPort := flag.String("dbp", "", "db port address")
	dbUser := flag.String("dbu", "", "db user")
	dbPass := flag.String("db_psw", "", "db pass")
	dbName := flag.String("dbn", "", "db name")
	dbSSL := flag.String("dbs", "", "db ssl mode")

	flag.Parse()

	return &AuthConfig{
		AuthAddr: *authAddr,
		AuthPort: *authPort,
		DBConfig: repository.DBConfig{
			Host:     *dbHost,
			Port:     *dbPort,
			Username: *dbUser,
			Password: *dbPass,
			DBName:   *dbName,
			SSLMode:  *dbSSL,
		},
	}
}

func initMigration(c repository.DBConfig) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Username, c.Password, c.Host, c.Port, c.DBName)
	m, err := migrate.New("file://db/migrations", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func initLog() {
	// log outputs
	file, err := os.OpenFile(LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
}
