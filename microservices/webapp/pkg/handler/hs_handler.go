package handler

import (
	"context"
	"github.com/askalag/aska/microservices/webapp/pkg/service"
	status_v1 "github.com/askalag/aska/microservices/webapp/proto/status/v1"
	"github.com/gin-gonic/gin"
)

type HSHandler struct {
	s service.History
}

func (h *HSHandler) status(ctx *gin.Context) {
	res, err := h.s.GrpcStatus(context.Background(), &status_v1.StatusRequest{})
	if err != nil {
		ctx.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"response": res,
	})

}

func NewHSHandler(s service.History) *HSHandler {
	return &HSHandler{s: s}
}
