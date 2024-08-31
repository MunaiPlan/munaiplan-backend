package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type HolesRepository interface {
	GetHoles(ctx context.Context, caseID string) ([]*entities.Hole, error)
	GetHoleByID(ctx context.Context, id string) (*entities.Hole, error)
	CreateHole(ctx context.Context, caseID string, hole *entities.Hole) error
	UpdateHole(ctx context.Context, hole *entities.Hole) (*entities.Hole, error)
	DeleteHole(ctx context.Context, id string) error
}