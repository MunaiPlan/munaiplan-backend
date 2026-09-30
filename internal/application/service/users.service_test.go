package service

import (
	"context"
	"errors"
	"testing"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/types"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
)

type fakeUsersRepo struct {
	user *entities.User
	err  error
}

func (f fakeUsersRepo) Create(context.Context, string, *entities.User) error { return nil }
func (f fakeUsersRepo) GetByEmail(context.Context, string) (*entities.User, error) {
	return f.user, f.err
}
func (f fakeUsersRepo) GetByID(context.Context, string) (*entities.User, error) {
	return f.user, f.err
}

func newTestJwt(t *testing.T) helpers.Jwt {
	t.Setenv("USER_ACCESS_TOKEN_SECRET", "test-access-secret")
	t.Setenv("USER_REFRESH_TOKEN_SECRET", "test-refresh-secret")
	t.Setenv("ACCESS_TOKEN_LIFETIME_MINUTES", "5")
	t.Setenv("REFRESH_TOKEN_LIFETIME_MINUTES", "10")
	jwt, err := helpers.NewJwt()
	if err != nil {
		t.Fatalf("jwt: %v", err)
	}
	return jwt
}

func TestSignInDoesNotRevealWhetherEmailExists(t *testing.T) {
	hash, err := helpers.HashPassword("correct-password")
	if err != nil {
		t.Fatal(err)
	}
	existing := &entities.User{ID: "u1", OrganizationID: "o1", Password: hash}
	dbDown := errors.New("connection refused")
	jwt := newTestJwt(t)

	cases := []struct {
		name     string
		repo     fakeUsersRepo
		password string
		wantErr  error
	}{
		{"unknown email", fakeUsersRepo{err: domainErrors.ErrUserNotFound}, "any-password", domainErrors.ErrInvalidCredentials},
		{"wrong password", fakeUsersRepo{user: existing}, "wrong-password", domainErrors.ErrInvalidCredentials},
		{"database failure", fakeUsersRepo{err: dbDown}, "correct-password", dbDown},
		{"success", fakeUsersRepo{user: existing}, "correct-password", nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewUsersService(tc.repo, nil, jwt)
			res, err := svc.SignIn(context.Background(), &requests.UserSignInRequest{Email: "a@b.local", Password: tc.password})
			if !errors.Is(err, tc.wantErr) || (tc.wantErr == nil && err != nil) {
				t.Fatalf("got error %v, want %v", err, tc.wantErr)
			}
			if tc.wantErr == nil && (res == nil || res.Token == "") {
				t.Fatalf("expected a token on success")
			}
		})
	}
}
