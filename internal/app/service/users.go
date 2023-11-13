package service

import (
	"context"
	"errors"
	"time"

	"github.com/munaiplan/munaiplan-backend/internal/app/domain"
	"github.com/munaiplan/munaiplan-backend/internal/app/repository"
	"github.com/munaiplan/munaiplan-backend/pkg/auth"
	"github.com/munaiplan/munaiplan-backend/pkg/hash"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersService struct {
	repo         repository.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	accessTokenTTL         time.Duration
	refreshTokenTTL        time.Duration
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher, accessTTL, refreshTTL time.Duration, 
	tokenManager auth.TokenManager) *UsersService {
	return &UsersService{
		repo:                   repo,
		hasher:                 hasher,
		tokenManager:           tokenManager,
		accessTokenTTL:         accessTTL,
		refreshTokenTTL:        refreshTTL,
	}
}

func (s *UsersService) SignUp(ctx context.Context, input UserSignUpInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	//verificationCode := s.otpGenerator.RandomSecret(s.verificationCodeLength)

	user := domain.User{
		Name:         input.Name,
		Password:     passwordHash,
		Phone:        input.Phone,
		Email:        input.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return err
		}

		return err
	}

	// return s.emailService.SendUserVerificationEmail(VerificationEmailInput{
	// 	Email:            user.Email,
	// 	Name:             user.Name,
	// 	VerificationCode: verificationCode,
	// })
	return err
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return Tokens{}, err
	}

	user, err := s.repo.GetByCredentials(ctx, input.Email, passwordHash)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return Tokens{}, err
		}

		return Tokens{}, err
	}

	return s.createSession(ctx, user.ID)
}

func (s *UsersService) createSession(ctx context.Context, userId primitive.ObjectID) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(userId.Hex(), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}

	err = s.repo.SetSession(ctx, userId, session)

	return res, err
}