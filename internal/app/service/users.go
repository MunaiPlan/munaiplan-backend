package service

import (
	"context"
	"errors"
	"time"

	"github.com/munaiplan/munaiplan-backend/internal/app/domain"
	"github.com/munaiplan/munaiplan-backend/internal/app/repository"
	"github.com/munaiplan/munaiplan-backend/pkg/hash"
)

type UsersService struct {
	repo         repository.Users
	hasher       hash.PasswordHasher
	//tokenManager auth.TokenManager
	accessTokenTTL         time.Duration
	refreshTokenTTL        time.Duration
	domain string
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher, accessTTL, refreshTTL time.Duration, domain string) *UsersService {
	return &UsersService{
		repo:                   repo,
		hasher:                 hasher,
		//tokenManager:           tokenManager,
		accessTokenTTL:         accessTTL,
		refreshTokenTTL:        refreshTTL,
		domain:                 domain,
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

	// todo. DECIDE ON EMAIL MARKETING STRATEGY

	// return s.emailService.SendUserVerificationEmail(VerificationEmailInput{
	// 	Email:            user.Email,
	// 	Name:             user.Name,
	// 	VerificationCode: verificationCode,
	// })
	return err
}