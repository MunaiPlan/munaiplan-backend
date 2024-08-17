package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/application/service"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/types"
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
		h.initFieldsRoutes(v1)
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

func (h *Handler) validateQueryIDParam(c *gin.Context, key string) (string, error) {
	value := c.Query(key)
	if value == "" {
		helpers.NewErrorResponse(c, http.StatusBadRequest, key + " is required")
		return "", errors.New(key + " is required")
	}
	if err := h.validateUUIDParam(c, value); err != nil {
		return "", err
	}
	return value, nil
}

func (h *Handler) validateRequestParam(c *gin.Context, key string) (string, error) {
	value := c.Param(key)
	if value == "" {
		helpers.NewErrorResponse(c, http.StatusBadRequest, key + " is required")
		return "", errors.New(key + " is required")
	}
	return value, nil
}

func (h *Handler) validateRequestIDParam(c *gin.Context, key string) (string, error) {
	value := c.Param(key)
	if value == "" {
		helpers.NewErrorResponse(c, http.StatusBadRequest, key + " is required")
		return "", errors.New(key + " is required")
	}
	if err := h.validateUUIDParam(c, value); err != nil {
		return "", err
	}
	return value, nil
}

func (h *Handler) validateContextIDKey(c *gin.Context, key string) (string, error) {
	value, exists := c.Get(key)
	if !exists {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, key + " is required")
		return "", errors.New(key + " is required")
	}
	valueString := value.(string)

	if err := h.validateUUIDParam(c, valueString); err != nil {
		return "", err
	}
	return valueString, nil
}

func (h *Handler) validateUUIDParam(c *gin.Context, value string) error {
	if err := uuid.Validate(value); err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, types.ErrInvalidUUID.Error())
		return errors.New(types.ErrInvalidUUID.Error())
	}
	return nil
}
