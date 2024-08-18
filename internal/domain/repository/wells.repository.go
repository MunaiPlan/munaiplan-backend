package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type WellsRepository interface {
	CreateWell(ctx context.Context, siteID string, well *entities.Well) error
	GetWellByID(ctx context.Context, id string) (*entities.Well, error)
	GetWells(ctx context.Context, siteID string) ([]*entities.Well, error)
	UpdateWell(ctx context.Context, well *entities.Well) (*entities.Well, error)
	DeleteWell(ctx context.Context, id string) error
}
