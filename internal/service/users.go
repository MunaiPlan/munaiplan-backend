package service

import (
	"context"
	"fmt"

	"github.com/munaiplan/munaiplan-backend/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/domain"
	"github.com/munaiplan/munaiplan-backend/internal/repository"
	"github.com/munaiplan/munaiplan-backend/presentation/types/requests"
	"github.com/munaiplan/munaiplan-backend/presentation/types/responses"
	"github.com/sirupsen/logrus"
)

const (
	BEARER_TOKEN_TYPE = "Bearer"
)

type UsersService struct {
	repo repository.UserRepository
	jwt  helpers.Jwt
}

func NewUsersService(repo repository.UserRepository, jwt helpers.Jwt) *UsersService {
	return &UsersService{
		repo: repo,
		jwt:  jwt,
	}
}
func (s *UsersService) SignUp(ctx context.Context, input requests.UserSignUpRequest) error {
	// Check if user with the same email already exists
	_, err := s.repo.GetByEmail(ctx, input.Email)
	if err == nil {
		return domain.ErrUserAlreadyExists
	}

	// Hash the password
	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		logrus.Errorf("Error hashing password: %s", err)
		return nil
	}

	// Create the user
	user := domain.User{
		Email:    input.Email,
		Password: hashedPassword,
	}

	// Save the user to the repository
	err = s.repo.Create(ctx, &user)
	if err != nil {
		logrus.Errorf("Error creating user: %s", err)
		return nil
	}

	return nil
}
func (s *UsersService) SignIn(ctx context.Context, input requests.UserSignInRequest) (*responses.TokenResponse, error) {
	user, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	fmt.Println(user.Password)
	fmt.Println(input.Password)

	if !helpers.CheckPasswordHash(input.Password, user.Password) {
		return nil, domain.ErrUserPasswordIncorrect
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
