package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/dto/requests"
	"github.com/munaiplan/munaiplan-backend/internal/application/dto/responses"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	domainErrors "github.com/munaiplan/munaiplan-backend/internal/domain/errors"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/sirupsen/logrus"
)

const (
	BEARER_TOKEN_TYPE = "Bearer"
)

type usersService struct {
	repo repository.UsersRepository
	jwt  helpers.Jwt
}

func NewUsersService(repo repository.UsersRepository, jwt helpers.Jwt) *usersService {
	return &usersService{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *usersService) SignUp(ctx context.Context, organizationId string, input *requests.UserSignUpRequest) error {
	// Check if user with the same email already exists
	_, err := s.repo.GetByEmail(ctx, organizationId, input.Email)
	if err == nil {
		return domainErrors.ErrUserAlreadyExists
	}

	// Hash the password
	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		logrus.Errorf("Error hashing password: %s", err)
		return nil
	}

	// Create the user
	user := entities.User{
		Email:    input.Email,
		Password: hashedPassword,
		Name:     input.Name,
		Surname:  input.Surname,
		Phone:    input.Phone,
	}

	// Save the user to the repository
	err = s.repo.Create(ctx, organizationId, &user)
	if err != nil {
		logrus.Errorf("Error creating user: %s", err)
		return nil
	}

	return nil
}
func (s *usersService) SignIn(ctx context.Context, organizationId string, input *requests.UserSignInRequest) (*responses.TokenResponse, error) {
	user, err := s.repo.GetByEmail(ctx, organizationId, input.Email)
	if err != nil {
		return nil, err
	}

	if !helpers.CheckPasswordHash(input.Password, user.Password) {
		return nil, domainErrors.ErrUserPasswordIncorrect
	}

	token, err := s.jwt.CreateAccessToken(helpers.UserAccessTokenClaims{UserId: user.ID})
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
