package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type FractureGradientsRepository interface {
	GetFractureGradients(ctx context.Context, caseID string) ([]*entities.FractureGradient, error)
	GetFractureGradientByID(ctx context.Context, id string) (*entities.FractureGradient, error)
	CreateFractureGradient(ctx context.Context, caseID string, fractureGradient *entities.FractureGradient) error
	UpdateFractureGradient(ctx context.Context, fractureGradient *entities.FractureGradient) (*entities.FractureGradient, error)
	DeleteFractureGradient(ctx context.Context, id string) error
}
