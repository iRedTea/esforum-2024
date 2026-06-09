package handler

import (
	"esforum/docs"
	"esforum/pkg/repository"
	"esforum/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
	repo     *repository.Repository
}

func NewHandler(services *service.Service, repo *repository.Repository) *Handler {
	return &Handler{services: services, repo: repo}
}

//	@title			EasyStartup Forum
//	@version		1.0
//	@description	Forum service for EasyStartup.
//
//	@contact.name	API Support
//	@contact.email	red.tea.dev@gmail.com
//
//	@host		forum.easystartup.su
//	@BasePath	/api
//

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	docs.SwaggerInfo.BasePath = "/"

	guest := router.Group("/api/guest")
	{
		guest.GET("thread/:id", nil)
		guest.GET("thread/all", nil)
		guest.GET("thread/:id/posts", nil)
		guest.GET("post/:id", nil)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.POST("thread", h.create)
		api.DELETE("thread/:id", h.delete)

		api.POST("post", nil)
		api.DELETE("post/:id", nil)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
