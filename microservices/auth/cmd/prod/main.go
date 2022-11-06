package main

import (
	"context"
	"flag"
	"github.com/Askalag/aska/microservices/auth/pkg/provider"
	"github.com/Askalag/aska/microservices/auth/pkg/repository"
	"github.com/Askalag/aska/microservices/auth/pkg/server"
	"github.com/Askalag/aska/microservices/auth/pkg/service"
	av1 "github.com/Askalag/aska/microservices/auth/proto/auth/v1"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var dbDriver = "postgres"
var dbAuthSchema = "auth"

type AuthConfig struct {
	AuthAddr  string
	AuthPort  string
	JWTConfig provider.JWTConfig
	DBConfig  repository.DBConfig
}

const LogFile = "/tmp/auth_log.log"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer ctx.Done()

	// app config
	c := buildConfig()

	// logger
	initLog()

	// Unit of Work repositories
	repos := repository.NewRepo(&c.DBConfig)

	// db migrations
	initMigration(c.DBConfig)

	// AuthProvider
	prov := provider.NewJWTProvider(&c.JWTConfig, &repos.AuthRepo)

	// services
	sessionService := service.NewSessionService(&repos.SessionRepo)
	authService := service.NewAuthService(repos.AuthRepo, prov, sessionService)

	// Unit of Work services
	services := service.NewService(sessionService, authService)

	// Unit of Work servers
	servers := server.NewServer(services)

	// GRPCServer
	grpcServer := grpc.NewServer()

	// Register servers in GRPCServer
	av1.RegisterAuthServiceServer(grpcServer, servers.Auth)

	// START
	startApp(ctx, cancel, c, grpcServer)
}

func startApp(ctx context.Context, cfn context.CancelFunc, cfg *AuthConfig, grpcServer *grpc.Server) {

	go notifySysCall(ctx, cfn)
	go listenAndServe(grpcServer, cfg)

	<-ctx.Done()
	log.Println("grpc server shutting down...")
	grpcServer.GracefulStop()
	log.Println("app has shutdown")
}

func listenAndServe(grpc *grpc.Server, cfg *AuthConfig) {
	// Up Listener
	listener, err := net.Listen("tcp", "localhost:"+cfg.AuthPort)
	if err != nil {
		log.Fatalln(err)
	}
	// Serve GRPC
	if err := grpc.Serve(listener); err != nil {
		log.Fatalf("failer to serve: '%v'", err)
	}
}

func notifySysCall(ctx context.Context, cf context.CancelFunc) {
	gCtx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	defer stop()

	for {
		select {
		case <-gCtx.Done():
			log.Println("syscall received, initiate graceful shutdown")
			cf()
			return
		default:
			time.Sleep(500 * time.Millisecond)
			continue
		}
	}
}

func buildConfig() *AuthConfig {
	authAddr := flag.String("auth_a", "", "http sign_in server address")
	authPort := flag.String("auth_p", "", "http sign_in port address")
	authKey := string(obtainRSAKey("private_rsa.pem"))

	dbHost := flag.String("dbh", "", "db host address")
	dbPort := flag.String("dbp", "", "db port address")
	dbUser := flag.String("dbu", "", "db user")
	dbPass := flag.String("db_psw", "", "db pass")
	dbName := flag.String("dbn", "", "db name")
	dbSSL := flag.String("db_ssl", "disable", "db ssl mode")
	dbSchema := dbAuthSchema

	flag.Parse()

	return &AuthConfig{
		AuthAddr:  *authAddr,
		AuthPort:  *authPort,
		JWTConfig: *provider.BuildJWTConfig(authKey, 60),
		DBConfig: *repository.NewDBConfig(
			*dbHost,
			*dbPort,
			*dbUser,
			*dbPass,
			*dbName,
			*dbSSL,
			dbSchema,
		),
	}
}

// obtainRSAKey Read RSA private key (file) for JWTProvider
func obtainRSAKey(pathToFile string) []byte {
	bytes, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		log.Fatalf("error oppening rsa file: '%s'", pathToFile)
	}
	return bytes
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
