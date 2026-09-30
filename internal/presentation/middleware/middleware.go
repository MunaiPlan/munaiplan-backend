package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/types"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
	"github.com/sirupsen/logrus"
)

// RoleCtx holds the authenticated account's role, read from the database per request.
const RoleCtx = "role"

// AccountLookup resolves the live account behind a token.
type AccountLookup interface {
	GetByID(ctx context.Context, userID string) (*entities.User, error)
}

// OwnershipLookup resolves the organization that owns a hierarchy record.
type OwnershipLookup interface {
	OrganizationOf(ctx context.Context, resource, id string) (string, error)
}

type AuthMiddleware struct {
	Jwt      helpers.Jwt
	Accounts AccountLookup
	Owners   OwnershipLookup
}

func NewAuthMiddleware(jwt helpers.Jwt, accounts AccountLookup, owners OwnershipLookup) *AuthMiddleware {
	return &AuthMiddleware{Jwt: jwt, Accounts: accounts, Owners: owners}
}

// routeResources maps the first path segment after /api/v1/ to the resource its :id names.
var routeResources = map[string]string{
	"companies": "companies", "fields": "fields", "sites": "sites", "wells": "wells",
	"wellbores": "wellbores", "designs": "designs", "trajectories": "trajectories", "cases": "cases",
	"holes": "holes", "strings": "strings", "fluids": "fluids", "rigs": "rigs",
	"pore-pressures": "pore_pressures", "fracture-gradients": "fracture_gradients",
}

// parentParams maps parent-id query parameters to their resource.
var parentParams = map[string]string{
	values.CompanyIdQueryParam: "companies", values.FieldIdQueryParam: "fields", values.SiteIdQueryParam: "sites",
	values.WellIdQueryParam: "wells", values.WellboreIdQueryParam: "wellbores", values.DesignIdQueryParam: "designs",
	values.TrajectoryIdQueryParam: "trajectories", values.CaseIdQueryParam: "cases",
}

// authorizeOwnership rejects any request naming a hierarchy record (path :id or parent query
// parameter) outside the caller's organization. Missing and foreign records both yield 404.
func (m *AuthMiddleware) authorizeOwnership(c *gin.Context, organizationID string) bool {
	check := func(resource, id string) bool {
		if _, err := uuid.Parse(id); err != nil {
			helpers.NewErrorResponse(c, http.StatusBadRequest, "invalid id")
			return false
		}
		owner, err := m.Owners.OrganizationOf(c.Request.Context(), resource, id)
		if err != nil {
			logrus.Errorf("resolve owner of %s: %v", resource, err)
			helpers.NewErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return false
		}
		if owner != organizationID {
			helpers.NewErrorResponse(c, http.StatusNotFound, "not found")
			return false
		}
		return true
	}
	if id := c.Param("id"); id != "" {
		segments := strings.Split(strings.TrimPrefix(c.FullPath(), "/api/v1/"), "/")
		if resource, ok := routeResources[segments[0]]; ok && !check(resource, id) {
			return false
		}
	}
	for param, resource := range parentParams {
		if id := c.Query(param); id != "" && !check(resource, id) {
			return false
		}
	}
	return true
}

// RequireAdmin must follow UserIdentity.
func (m *AuthMiddleware) RequireAdmin(c *gin.Context) {
	if c.GetString(RoleCtx) != entities.RoleAdmin {
		helpers.NewErrorResponse(c, http.StatusForbidden, domainErrors.ErrForbidden.Error())
		return
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

	// A valid signature is not enough: the account must still exist in the same organization,
	// so deleted accounts lose access immediately rather than at token expiry.
	account, err := m.Accounts.GetByID(c.Request.Context(), userClaims.UserId)
	if errors.Is(err, domainErrors.ErrUserNotFound) || (err == nil && account.OrganizationID != userClaims.OrganizationId) {
		helpers.NewErrorResponse(c, http.StatusUnauthorized, "session is no longer valid")
		return
	}
	if err != nil {
		logrus.Errorf("resolve account: %v", err)
		helpers.NewErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.Set(values.UserIdCtx, userClaims.UserId)
	c.Set(values.OrganizationIdCtx, userClaims.OrganizationId)
	c.Set(RoleCtx, account.Role)
	if !m.authorizeOwnership(c, account.OrganizationID) {
		return
	}
	c.Set(values.UserRefreshTokenCtx, header)
}