package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/errors"
	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/application/dto/responses"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-in", h.SignIn)
		users.POST("/sign-up", h.SignUp)
	}
}

// @Summary User SignUp
// @Tags users-auth
// @Description user sign up
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Param input body signUpInput true "sign up info"
// @Success 201 {object} tokenResponse
// @Failure 400,500 {object} response
// @Failure default {object} response
// @Router /users/sign-up [post]
func (h *Handler) SignUp(c *gin.Context) {
	var inp requests.UserSignUpRequest
	if err := c.BindJSON(&inp); err != nil {
		responses.NewResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Users.SignUp(c.Request.Context(), inp)
	if err != nil {
		responses.NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary User SignIn
// @Tags users-auth
// @Description user sign in
// @ModuleID userSignIn
// @Accept  json
// @Produce  json
// @Param input body signInInput true "sign in info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-in [post]
func (h *Handler) SignIn(c *gin.Context) {
	var inp requests.UserSignInRequest
	if err := c.BindJSON(&inp); err != nil {
		responses.NewResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Users.SignIn(c.Request.Context(), inp)
	if err != nil {
		if errors.Is(err, domainErrors.ErrUserNotFound) {
			responses.NewResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		responses.NewResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, &res)
}