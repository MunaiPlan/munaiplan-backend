package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type HolesRepository interface {
	// GetHoles retrieves all holes associated with a specific case ID.
	GetHoles(ctx context.Context, caseID string) ([]*entities.Hole, error)
	
	// GetHoleByID retrieves a single hole by its ID.
	GetHoleByID(ctx context.Context, id string) (*entities.Hole, error)
	
	// CreateHole creates a new hole and its associated caisings in the database.
	CreateHole(ctx context.Context, caseID string, hole *entities.Hole) error
	
	// UpdateHole updates an existing hole and its associated caisings in the database.
	UpdateHole(ctx context.Context, hole *entities.Hole) (*entities.Hole, error)
	
	// DeleteHole deletes a hole and its associated caisings from the database by hole ID.
	DeleteHole(ctx context.Context, id string) error
}
