package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/types"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// initUsersRoutes initializes the user routes.
func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-in", h.signIn)
		users.POST("/sign-up", h.signUp)
	}
}

// signUp handles the user sign up request.
// @Summary User SignUp
// @Tags users-auth
// @Description user sign up
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Param organizationId query string true "Organization ID"
// @Param input body requests.UserSignUpRequest true "sign up info"
// @Success 201 {object} helpers.Response
// @Failure 400,500 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /api/v1/users/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var inp requests.UserSignUpRequest
	var err error
	if inp.OrganizationID, err = h.validateQueryIDParam(c, values.OrganizationIdQueryParam); err != nil {		
		return
	}

	if err := c.BindJSON(&inp.Body); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.services.Users.SignUp(c.Request.Context(), &inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("user created"))
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
// @Failure 400,404 {object} helpers.Response
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
		if errors.Is(err, domainErrors.ErrUserNotFound) {
			helpers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &res)
}