package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type UsersRepository interface {
    Create(ctx context.Context, organizationId string, user *entities.User) error
    GetByEmail(ctx context.Context, organizationId string, email string) (*entities.User, error)
    // Update(ctx context.Context, user domain.User) error
}