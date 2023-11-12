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

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) error
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
	Domain                 string
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.Hasher,
		deps.AccessTokenTTL, deps.RefreshTokenTTL, deps.Domain)

	return &Services{
		Users:    usersService,
	}
}