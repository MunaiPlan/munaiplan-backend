package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/service"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
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
		h.initOrganizationsRoutes(v1)
		h.initCompaniesRoutes(v1)
	}
}

func (h *Handler) validateQueryParam(c *gin.Context, key string) (string, error) {
	value := c.Query(key)
	if value == "" {
		helpers.NewErrorResponse(c, http.StatusBadRequest, key + " is required")
		return "", errors.New(key + " is required")
	}
	return value, nil
}