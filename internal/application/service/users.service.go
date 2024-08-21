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

func (s *usersService) SignUp(ctx context.Context, input *requests.UserSignUpRequest) error {
	if err := s.commonRepo.CheckIfUserExistsByEmail(ctx, input.Body.Email); err == nil {
		return domainErrors.ErrUserAlreadyExists
	}

	// Hash the password
	hashedPassword, err := helpers.HashPassword(input.Body.Password)
	if err != nil {
		logrus.Errorf("Error hashing password: %s", err)
		return nil
	}

	// Create the user
	user := entities.User{
		Email:    input.Body.Email,
		Password: hashedPassword,
		Name:     input.Body.Name,
		Surname:  input.Body.Surname,
		Phone:    input.Body.Phone,
		OrganizationID: input.OrganizationID,
	}

	// Save the user to the repository
	err = s.repo.Create(ctx, input.OrganizationID, &user)
	if err != nil {
		logrus.Errorf("Error creating user: %s", err)
		return err
	}

	return nil
}
func (s *usersService) SignIn(ctx context.Context, input *requests.UserSignInRequest) (*responses.TokenResponse, error) {
	if err := s.commonRepo.CheckIfUserExistsByEmail(ctx, input.Email); err != nil {
		return nil, err
	}

	user, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if !helpers.CheckPasswordHash(input.Password, user.Password) {
		return nil, domainErrors.ErrUserPasswordIncorrect
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
