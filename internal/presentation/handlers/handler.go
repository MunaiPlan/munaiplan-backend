package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/service"
	"github.com/munaiplan/munaiplan-backend/internal/presentation/middleware"
)

type Handler struct {
	services       *service.Services
	authMiddleware *middleware.AuthMiddleware
}

func NewHandler(services *service.Services, authMiddleware *middleware.AuthMiddleware) *Handler {
	return &Handler{
		services:       services,
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
		h.initSitesRoutes(v1)
		h.initWellsRoutes(v1)
		h.initWellboresRoutes(v1)
		h.initDesignsRoutes(v1)
		h.initTrajectoriesRoutes(v1)
		h.initCasesRoutes(v1)
		h.initHolesRoutes(v1)
		h.initFluidsRoutes(v1)
		h.initRigsRoutes(v1)
		h.initPorePressureRoutes(v1)
		h.initPressureDataProfilesRoutes(v1)
		h.initFractureGradientRoutes(v1)
		h.initStringsRoutes(v1)
	}
}
