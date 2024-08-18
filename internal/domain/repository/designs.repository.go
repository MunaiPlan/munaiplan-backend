package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type DesignsRepository interface {
	CreateDesign(ctx context.Context, wellboreID string, design *entities.Design) error
	GetDesignByID(ctx context.Context, id string) (*entities.Design, error)
	GetDesigns(ctx context.Context, wellboreID string) ([]*entities.Design, error)
	UpdateDesign(ctx context.Context, design *entities.Design) (*entities.Design, error)
	DeleteDesign(ctx context.Context, id string) error
}