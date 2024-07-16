package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/errors"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-in", h.signIn)
		users.POST("/sign-up", h.signUp)
	}
}

// @Summary User SignUp
// @Tags users-auth
// @Description user sign up
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Param input body requests.UserSignUpRequest true "sign up info"
// @Success 201 {object} helpers.Response
// @Failure 400,500 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /api/v1/users/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var inp *requests.UserSignUpRequest
	organizationId, err := h.validateQueryParam(c, values.OrganizationIdQueryParam)	
	if err != nil {
		return
	}

	if err := c.BindJSON(&inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	fmt.Println(inp.Phone + " edf")

	err = h.services.Users.SignUp(c.Request.Context(), organizationId, inp)
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, helpers.NewResponse("user created"))
}

// @Summary User SignIn
// @Tags users-auth
// @Description user sign in
// @ModuleID userSignIn
// @Accept  json
// @Produce  json
// @Param input body requests.UserSignInRequest true "sign in info"
// @Success 200 {object} helpers.TokenResponse
// @Failure 400,404 {object} helpers.Response
// @Failure 500 {object} helpers.Response
// @Failure default {object} helpers.Response
// @Router /api/v1/users/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var inp *requests.UserSignInRequest
	organizationId, err := h.validateQueryParam(c, values.OrganizationIdQueryParam)
	if err != nil {
		return
	}

	if err := c.BindJSON(&inp); err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Users.SignIn(c.Request.Context(), organizationId, inp)
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