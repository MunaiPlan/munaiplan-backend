package service

import (
	"context"
	"errors"
	"sync"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/responses"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/types"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/sirupsen/logrus"
)

const (
	BEARER_TOKEN_TYPE = "Bearer"
)

type usersService struct {
	repo repository.UsersRepository
	commonRepo repository.CommonRepository
	jwt  helpers.Jwt
}

func NewUsersService(repo repository.UsersRepository, commonRepo repository.CommonRepository, jwt helpers.Jwt) *usersService {
	return &usersService{
		repo: repo,
		commonRepo: commonRepo,
		jwt:  jwt,
	}
}

func (s *usersService) GetByID(ctx context.Context, userID string) (*entities.User, error) {
	return s.repo.GetByID(ctx, userID)
}

// dummyPasswordHash lets an unknown email cost the same bcrypt work as a real
// account, so response time does not reveal which emails exist.
var dummyPasswordHash = sync.OnceValue(func() string {
	hash, _ := helpers.HashPassword("unused-timing-equalizer")
	return hash
})

func (s *usersService) SignIn(ctx context.Context, input *requests.UserSignInRequest) (*responses.TokenResponse, error) {
	user, err := s.repo.GetByEmail(ctx, normalizeEmail(input.Email))
	if errors.Is(err, domainErrors.ErrUserNotFound) {
		helpers.CheckPasswordHash(input.Password, dummyPasswordHash())
		return nil, domainErrors.ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}

	if !helpers.CheckPasswordHash(input.Password, user.Password) {
		return nil, domainErrors.ErrInvalidCredentials
	}

	token, err := s.jwt.CreateAccessToken(helpers.UserAccessTokenClaims{
		UserId: user.ID,
		OrganizationId: user.OrganizationID,
	})
	if err != nil {
		logrus.Errorf("Error creating access token: %s", err)
		return nil, err
	}
	return &responses.TokenResponse{
		Success:               true,
		Token:                 token.AccessToken,
		TokenType:             BEARER_TOKEN_TYPE,
		ExpiresAt:             token.AccessTokenExpiresAt,
		RefreshToken:          token.RefreshToken,
		RefreshTokenType:      BEARER_TOKEN_TYPE,
		RefreshTokenExpiresAt: token.RefreshTokenExpiresAt,
	}, nil
}
