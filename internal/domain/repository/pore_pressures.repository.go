package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type PorePressuresRepository interface {
	CreatePorePressure(ctx context.Context, caseID string, porePressure *entities.PorePressure) error
	GetPorePressureByID(ctx context.Context, id string) (*entities.PorePressure, error)
	GetPorePressures(ctx context.Context, caseID string) ([]*entities.PorePressure, error)
	UpdatePorePressure(ctx context.Context, porePressure *entities.PorePressure) (*entities.PorePressure, error)
	DeletePorePressure(ctx context.Context, id string) error
}
