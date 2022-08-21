package main

import (
	"flag"
	"fmt"
	"github.com/askalag/aska/microservices/webapp/pkg"
	"github.com/askalag/aska/microservices/webapp/pkg/handler"
	"github.com/askalag/aska/microservices/webapp/pkg/logging"
	"github.com/askalag/aska/microservices/webapp/pkg/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {

	config := buildConfig()

	logging.Initial(config.Log)
	logrus.Infoln("logger initialized")

	services := service.NewService(config.ServicesTCP)
	handlers := handler.NewHandler(services)
	engine := handler.NewEngine(handlers)

	startApp(engine, config)

}

func startApp(h http.Handler, c pkg.Config) {
	// http server start up
	logrus.Infoln("Listen and serve on:", c.ServicesTCP.AppAddr)
	server := pkg.NewServer(h, c)
	err := server.ListenAndServe()
	if err != nil {
		logrus.Fatalln(err.Error())
	}
}

func buildConfig() pkg.Config {
	// get params from command line
	addr := flag.String("app_a", "", "http server address")
	port := flag.String("app_p", "", "http server port")
	authAddr := flag.String("auth_a", "", "http auth server address")
	authPort := flag.String("auth_p", "", "http auth port address")
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

	sLOG := pkg.Logging{
		Format:   *logFormat,
		Level:    *logLevel,
		FilePath: "/tmp/webapi_log.log",
	}

	return pkg.Config{
		ServicesTCP:  sTCP,
		IdleTimeout:  1 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Log:          sLOG,
	}
}
