package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

type AuthMiddleware struct {
	Jwt               helpers.Jwt
}

func NewAuthMiddleware(jwt helpers.Jwt) *AuthMiddleware {
	return &AuthMiddleware{
		Jwt:               jwt,
	}
}

func (m *AuthMiddleware) RefreshTokenIdentity(c *gin.Context) {
	header := c.GetHeader(values.AuthorizationHeader)
	if header == "" {
		helpers.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		helpers.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userClaims, err := m.Jwt.VerifyRefreshToken(headerParts[1])
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(values.UserIdCtx, userClaims.UserId)
	c.Set(values.UserRefreshTokenCtx, header)
}

func (m *AuthMiddleware) UserIdentity(c *gin.Context) {
	header := c.GetHeader(values.AuthorizationHeader)
	if header == "" {
		helpers.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		helpers.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userClaims, err := m.Jwt.Verify(headerParts[1])
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(values.UserIdCtx, userClaims.UserId)
	c.Set(values.UserRefreshTokenCtx, header)
}