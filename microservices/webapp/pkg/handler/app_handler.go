package handler

import (
	"github.com/askalag/aska/microservices/webapp/pkg/service"
	"github.com/gin-gonic/gin"
)

type AppHandler struct {
	s service.App
}

func (h *AppHandler) status(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "alive",
	})
}

func NewAppHandler(s service.App) *AppHandler {
	return &AppHandler{s: s}
}
