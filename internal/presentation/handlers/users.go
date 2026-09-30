package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/types"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
	"github.com/sirupsen/logrus"
)

// initUsersRoutes initializes the user routes.
func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-in", h.signIn)
		users.GET("/me", h.authMiddleware.UserIdentity, h.me)
	}
	api.GET("/status", h.authMiddleware.UserIdentity, h.status)
}

type meResponse struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	Role           string `json:"role"`
}

// status reports service health for the UI status bar. "api" is always ok when this answers.
// @Summary Service status
// @Tags users-auth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} map[string]string
// @Router /api/v1/status [get]
func (h *Handler) status(c *gin.Context) {
	model := "ready"
	if err := h.services.TorqueAndDrag.ModelReady(c.Request.Context()); err != nil {
		model = "unavailable"
	}
	c.JSON(http.StatusOK, gin.H{"api": "ok", "model": model})
}

// me returns the signed-in account, including its role for the admin panel.
// @Summary Current user
// @Tags users-auth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} meResponse
// @Failure 401 {object} helpers.Response
// @Router /api/v1/users/me [get]
func (h *Handler) me(c *gin.Context) {
	user, err := h.services.Users.GetByID(c.Request.Context(), c.GetString(values.UserIdCtx))
	if err != nil {
		h.respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, meResponse{ID: user.ID, OrganizationID: user.OrganizationID, Name: user.Name, Surname: user.Surname, Email: user.Email, Role: user.Role})
}

// signIn handles the user sign in request.
// @Summary User SignIn
// @Tags users-auth
// @Description user sign in
// @ModuleID userSignIn
// @Accept  json
// @Produce  json
// @Param organizationId query string true "Organization ID"
// @Param input body requests.UserSignInRequest true "sign in info"
// @Success 200 {object} responses.TokenResponse
// @Failure 400,401 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /api/v1/users/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var inp requests.UserSignInRequest
	if err := c.BindJSON(&inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Users.SignIn(c.Request.Context(), &inp)
	if err != nil {
		if errors.Is(err, domainErrors.ErrInvalidCredentials) {
			helpers.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		logrus.Errorf("sign-in failed: %v", err)
		helpers.NewErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, &res)
}