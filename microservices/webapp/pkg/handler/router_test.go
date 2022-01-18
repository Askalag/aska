package handler_test

import (
	"github.com/askalag/aska/microservices/webapp/pkg"
	"github.com/askalag/aska/microservices/webapp/pkg/handler"
	"github.com/askalag/aska/microservices/webapp/pkg/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	services := service.NewService(pkg.ServicesTCP{})
	handlers := handler.NewHandler(services)
	return handler.NewEngine(handlers)
}
