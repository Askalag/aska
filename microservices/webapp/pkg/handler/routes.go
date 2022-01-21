package handler

import (
	"github.com/askalag/aska/microservices/webapp/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	app *AppHandler
	hs  *HSHandler
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		app: NewAppHandler(s.App),
		hs:  NewHSHandler(s.History),
	}
}

func NewEngine(h *Handler) *gin.Engine {
	r := gin.New()

	api := r.Group("/api")
	{
		//sign_in := api.Group("/sign_in")
		app := api.Group("/app")
		{
			app.GET("status", h.app.status)
		}

		hs := api.Group("/hs")
		{
			hs.GET("status", h.hs.status)
		}
	}
	return r
}
