package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/app/domain"
	"github.com/munaiplan/munaiplan-backend/internal/app/service"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignUp)
	}
}

type userSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Phone    string `json:"phone" binding:"required,max=13"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

// @Summary User SignUp
// @Tags users-auth
// @Description create user account
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Param input body userSignUpInput true "sign up info"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-up [post]
func (h *Handler) userSignUp(c *gin.Context) {
	var inp userSignUpInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	if err := h.services.Users.SignUp(c.Request.Context(), service.UserSignUpInput{
		Name:     inp.Name,
		Email:    inp.Email,
		Phone:    inp.Phone,
		Password: inp.Password,
	}); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusCreated)
}