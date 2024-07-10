package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/service"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/middleware"
)

type Handler struct {
	services *service.Services
	authMiddleware *middleware.AuthMiddleware
}

func NewHandler(services *service.Services, authMiddleware *middleware.AuthMiddleware) *Handler {
	return &Handler{
		services: services,
		authMiddleware: authMiddleware,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)
	}
}
