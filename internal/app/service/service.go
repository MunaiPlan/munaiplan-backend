package service

import (
	"context"
	"time"

	"github.com/munaiplan/munaiplan-backend/internal/app/repository"
	"github.com/munaiplan/munaiplan-backend/pkg/auth"
	"github.com/munaiplan/munaiplan-backend/pkg/hash"
)

type UserSignUpInput struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

type UserSignInInput struct {
	Email    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (Tokens, error)
}

type Services struct {
	Users          Users
}

type Deps struct {
	Repos                  *repository.Repositories
	Hasher                 hash.PasswordHasher
	TokenManager           auth.TokenManager
	AccessTokenTTL         time.Duration
	RefreshTokenTTL        time.Duration
	Environment            string
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.Hasher,
		deps.AccessTokenTTL, deps.RefreshTokenTTL, deps.TokenManager)

	return &Services{
		Users:    usersService,
	}
}