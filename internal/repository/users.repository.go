package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain"
)

type UserRepository interface {
    Create(ctx context.Context, user *domain.User) error
    GetByEmail(ctx context.Context, email string) (*domain.User, error)
    // Update(ctx context.Context, user domain.User) error
}