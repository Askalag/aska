package main

import (
	"flag"
	"fmt"
	"github.com/askalag/aska/microservices/webapp/pkg"
	"github.com/askalag/aska/microservices/webapp/pkg/handler"
	"github.com/askalag/aska/microservices/webapp/pkg/service"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

const LogFile = "/tmp/webapi_log.log"

func main() {

	config := buildConfig()
	services := service.NewService(config.ServicesTCP)
	handlers := handler.NewHandler(services)
	engine := handler.NewEngine(handlers)

	startApp(engine, config)

}

func startApp(h http.Handler, c *pkg.Config) {

	iniLogger(c.LogFmt, c.LogLevel)

	// http server start up
	log.Infoln("Listen and serve on:", c.ServicesTCP.AppAddr)
	server := pkg.NewServer(h, c)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func buildConfig() *pkg.Config {
	// get params from command line
	addr := flag.String("app_a", "", "http server address")
	port := flag.String("app_p", "", "http server port")
	authAddr := flag.String("auth_a", "", "http sign_in server address")
	authPort := flag.String("auth_p", "", "http sign_in port address")
	taskAddr := flag.String("task_a", "", "http task server address")
	taskPort := flag.String("task_p", "", "http task port address")
	hsAddr := flag.String("hs_a", "", "http history server address")
	hsPort := flag.String("hs_p", "", "http history port address")
	logFormat := flag.String("logf", "text", "set log format")
	logLevel := flag.String("logl", "", "log level")
	flag.Parse()

	sTCP := pkg.ServicesTCP{
		AppAddr:     fmt.Sprintf("%s:%s", *addr, *port),
		AuthAddr:    fmt.Sprintf("%s:%s", *authAddr, *authPort),
		HistoryAddr: fmt.Sprintf("%s:%s", *hsAddr, *hsPort),
		TaskAddr:    fmt.Sprintf("%s:%s", *taskAddr, *taskPort),
	}

	return &pkg.Config{
		ServicesTCP:  sTCP,
		IdleTimeout:  1 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		LogFmt:       *logFormat,
		LogLevel:     *logLevel,
	}
}

func iniLogger(fmt string, lv string) {
	// log format
	switch fmt {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
		log.Infoln("Logger format: json")
	case "text":
		log.SetFormatter(&log.TextFormatter{})
		log.Infoln("Logger format: text")
	default:
		log.Warningln("cannot set log format, use default 'text'")
	}

	// log level
	level, err := log.ParseLevel(lv)
	if err != nil {
		log.Errorln("cannot set log level, use default 'debug'")
		level = log.DebugLevel
	} else {
		log.Infoln("Logger mode level:", level.String())
	}
	log.SetLevel(level)

	// log outputs
	file, err := os.OpenFile(LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
}
