package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type WellboresRepository interface {
	CreateWellbore(ctx context.Context, wellID string, wellbore *entities.Wellbore) error
	GetWellboreByID(ctx context.Context, id string) (*entities.Wellbore, error)
	GetWellbores(ctx context.Context, wellID string) ([]*entities.Wellbore, error)
	UpdateWellbore(ctx context.Context, wellbore *entities.Wellbore) (*entities.Wellbore, error)
	DeleteWellbore(ctx context.Context, id string) error
}
