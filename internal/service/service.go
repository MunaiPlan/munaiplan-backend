package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/repository"
	"github.com/munaiplan/munaiplan-backend/presentation/types/requests"
	"github.com/munaiplan/munaiplan-backend/presentation/types/responses"
)

type Users interface {
	SignIn(ctx context.Context, input requests.UserSignInRequest) (*responses.TokenResponse, error)
	SignUp(ctx context.Context, input requests.UserSignUpRequest) error
}

type Services struct {
	// TODO() Implement cache
	// CatalogCache *catalog.CatalogCache
	Users
}


func NewServices(repos *repository.Repository, jwt helpers.Jwt) *Services {
	usersService := NewUsersService(repos.Users, jwt)

	return &Services{
		Users:        usersService,
		// CatalogCache: deps.CatalogCache,
	}
}
