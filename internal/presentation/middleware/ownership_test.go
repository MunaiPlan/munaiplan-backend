package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type fakeOwners map[string]string // id -> organization

func (f fakeOwners) OrganizationOf(_ context.Context, _ string, id string) (string, error) {
	return f[id], nil
}

const (
	mine    = "11111111-1111-1111-1111-111111111111"
	foreign = "22222222-2222-2222-2222-222222222222"
	missing = "33333333-3333-3333-3333-333333333333"
	orgA    = "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
	orgB    = "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"
)

func TestAuthorizeOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)
	m := &AuthMiddleware{Owners: fakeOwners{mine: orgA, foreign: orgB}}
	router := gin.New()
	guard := func(c *gin.Context) {
		if m.authorizeOwnership(c, orgA) {
			c.Status(http.StatusOK)
		}
	}
	router.GET("/api/v1/cases/:id", guard)
	router.GET("/api/v1/strings/", guard)
	router.GET("/api/v1/fluids/types", guard)

	cases := map[string]int{
		"/api/v1/cases/" + mine:              200,
		"/api/v1/cases/" + foreign:           404,
		"/api/v1/cases/" + missing:           404,
		"/api/v1/cases/not-a-uuid":           400,
		"/api/v1/strings/?caseId=" + mine:    200,
		"/api/v1/strings/?caseId=" + foreign: 404,
		"/api/v1/strings/":                   200,
		"/api/v1/fluids/types":               200,
	}
	for path, want := range cases {
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
		if rec.Code != want {
			t.Errorf("%s: got %d, want %d", path, rec.Code, want)
		}
	}
}
