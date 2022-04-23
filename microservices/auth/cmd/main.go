package main

import (
	"flag"
	"github.com/Askalag/aska/microservices/auth/internal/provider"
	"github.com/Askalag/aska/microservices/auth/internal/repository"
	"github.com/Askalag/aska/microservices/auth/internal/server"
	"github.com/Askalag/aska/microservices/auth/internal/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

var dbDriver = "postgres"

type AuthConfig struct {
	AuthAddr  string
	AuthPort  string
	JWTConfig provider.JWTConfig
	DBConfig  repository.DBConfig
}

const LogFile = "/tmp/auth_log.log"

func main() {
	// global config
	c := buildConfig()

	// logger
	initLog()

	// db migrations
	initMigration(c.DBConfig)

	// init repo, service, jwt providers, servers
	repos := repository.NewRepo(&c.DBConfig)
	prov := provider.NewJWTProvider(&c.JWTConfig, &repos.AuthRepo)

	services := service.NewService(repos, prov)
	_ = server.NewServer(services)

}

func buildConfig() *AuthConfig {
	authAddr := flag.String("auth_a", "", "http sign_in server address")
	authPort := flag.String("auth_p", "", "http sign_in port address")
	authKey := flag.String("auth_k", "", "auth security key")

	dbHost := flag.String("dbh", "", "db host address")
	dbPort := flag.String("dbp", "", "db port address")
	dbUser := flag.String("dbu", "", "db user")
	dbPass := flag.String("db_psw", "", "db pass")
	dbName := flag.String("dbn", "", "db name")
	dbSSL := flag.String("db_ssl", "disable", "db ssl mode")

	flag.Parse()

	return &AuthConfig{
		AuthAddr:  *authAddr,
		AuthPort:  *authPort,
		JWTConfig: *provider.BuildJWTConfig(*authKey, 60),
		DBConfig: *repository.NewDBConfig(
			*dbHost,
			*dbPort,
			*dbUser,
			*dbPass,
			*dbName,
			*dbSSL,
		),
	}
}

func initMigration(c repository.DBConfig) {
	dbURL := c.BuildConnString(dbDriver)
	m, err := migrate.New("file://db/migrations", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func initLog() {
	// log outputs file and print
	file, err := os.OpenFile(LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
}
