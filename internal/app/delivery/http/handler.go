package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
	"github.com/munaiplan/munaiplan-backend/internal/app/config"
	_ "github.com/munaiplan/munaiplan-backend/docs"

	v1 "github.com/munaiplan/munaiplan-backend/internal/app/delivery/http/v1"
	"github.com/munaiplan/munaiplan-backend/internal/app/service"
	"github.com/munaiplan/munaiplan-backend/pkg/auth"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	// Init gin handler
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)


	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}