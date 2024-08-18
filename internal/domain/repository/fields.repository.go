package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type FieldsRepository interface {
	CreateField(ctx context.Context, companyID string, field *entities.Field) error
	GetFieldByID(ctx context.Context, id string) (*entities.Field, error)
	GetFields(ctx context.Context, companyID string) ([]*entities.Field, error)
	UpdateField(ctx context.Context, field *entities.Field) (*entities.Field, error)
	DeleteField(ctx context.Context, id string) error
}
