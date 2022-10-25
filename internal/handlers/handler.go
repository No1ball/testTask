package handlers

import (
	"github.com/No1ball/testEx/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		followers := api.Group("/followers")
		{
			followers.GET("/", h.GetFollowers)
			followers.GET("/:id", h.GetFollowerById)
			followers.POST("/:id", h.AddFollower)
		}

		templates := api.Group("/templates")
		{
			templates.GET("/", h.GetFollowers)
			templates.GET("/:id", h.GetTemplateById)
			templates.POST("/:id", h.AddTemplate)
		}
		api.GET("/email", h.SendEmail)
	}

	return router
}
