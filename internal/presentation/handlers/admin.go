package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/service"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/types"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	client "github.com/munaiplan/munaiplan-backend/internal/infrastructure/prediction_client"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// initAdminRoutes exposes tenant and account provisioning to administrators only.
func (h *Handler) initAdminRoutes(api *gin.RouterGroup) {
	admin := api.Group("/admin", h.authMiddleware.UserIdentity, h.authMiddleware.RequireAdmin)
	{
		admin.GET("/organizations", h.adminListOrganizations)
		admin.POST("/organizations", h.adminCreateOrganization)
		admin.GET("/organizations/:id/users", h.adminListUsers)
		admin.POST("/organizations/:id/users", h.adminCreateUser)
	}
}

// @Summary List organizations (admin)
// @Tags admin
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} entities.OrganizationSummary
// @Router /api/v1/admin/organizations [get]
func (h *Handler) adminListOrganizations(c *gin.Context) {
	orgs, err := h.services.Admin.ListOrganizations(c.Request.Context())
	if err != nil {
		h.respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, orgs)
}

// @Summary Create an organization with its first user (admin)
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param input body requests.AdminCreateOrganizationRequest true "organization and first user"
// @Success 201 {object} map[string]interface{}
// @Failure 400,409 {object} helpers.Response
// @Router /api/v1/admin/organizations [post]
func (h *Handler) adminCreateOrganization(c *gin.Context) {
	var in requests.AdminCreateOrganizationRequest
	if err := c.ShouldBindJSON(&in); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid input body: "+err.Error())
		return
	}
	org, user, err := h.services.Admin.CreateOrganization(c.Request.Context(), &in)
	if err != nil {
		h.respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"organization": org, "user": user})
}

// @Summary List an organization's users (admin)
// @Tags admin
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Organization ID"
// @Success 200 {array} entities.User
// @Router /api/v1/admin/organizations/{id}/users [get]
func (h *Handler) adminListUsers(c *gin.Context) {
	users, err := h.services.Admin.ListUsers(c.Request.Context(), c.Param("id"))
	if err != nil {
		h.respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Add a user to an organization (admin)
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Organization ID"
// @Param input body requests.AdminUserInput true "user"
// @Success 201 {object} entities.User
// @Failure 400,404,409 {object} helpers.Response
// @Router /api/v1/admin/organizations/{id}/users [post]
func (h *Handler) adminCreateUser(c *gin.Context) {
	var in requests.AdminUserInput
	if err := c.ShouldBindJSON(&in); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid input body: "+err.Error())
		return
	}
	user, err := h.services.Admin.CreateUser(c.Request.Context(), c.Param("id"), &in)
	if err != nil {
		h.respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, user)
}

// respondError maps domain errors to HTTP statuses and hides internal details.
func (h *Handler) respondError(c *gin.Context, err error) {
	var validation *service.ValidationError
	var duplicate *service.ErrAlreadyImported
	var caseData *service.CaseDataError
	switch {
	case errors.As(err, &caseData):
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error(), "problems": caseData.Problems})
	case errors.Is(err, client.ErrUnavailable):
		logrus.Warnf("%s %s: %v", c.Request.Method, c.FullPath(), err)
		helpers.NewErrorResponse(c, http.StatusServiceUnavailable, client.ErrUnavailable.Error())
	case errors.Is(err, client.ErrInvalidResponse):
		logrus.Errorf("%s %s: %v", c.Request.Method, c.FullPath(), err)
		helpers.NewErrorResponse(c, http.StatusBadGateway, client.ErrInvalidResponse.Error())
	case errors.Is(err, gorm.ErrRecordNotFound):
		helpers.NewErrorResponse(c, http.StatusNotFound, "not found")
	case errors.As(err, &validation):
		helpers.NewErrorResponse(c, http.StatusBadRequest, validation.Reason)
	case errors.As(err, &duplicate):
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error(), "case_id": duplicate.CaseID})
	case errors.Is(err, domainErrors.ErrEmailTaken):
		helpers.NewErrorResponse(c, http.StatusConflict, err.Error())
	case errors.Is(err, domainErrors.ErrOrganizationNotFound), errors.Is(err, domainErrors.ErrUserNotFound):
		helpers.NewErrorResponse(c, http.StatusNotFound, err.Error())
	default:
		logrus.Errorf("%s %s: %v", c.Request.Method, c.FullPath(), err)
		helpers.NewErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
}
