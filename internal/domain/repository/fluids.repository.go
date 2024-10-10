package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

// FluidsRepository defines the interface for fluid-related database operations.
type FluidsRepository interface {
	// GetFluids retrieves all fluids associated with a specific case.
	GetFluids(ctx context.Context, caseID string) ([]*entities.Fluid, error)
	
	// GetFluidByID retrieves a fluid by its ID.
	GetFluidByID(ctx context.Context, id string) (*entities.Fluid, error)
	
	// CreateFluid creates a new fluid within a specified case.
	CreateFluid(ctx context.Context, caseID string, fluid *entities.Fluid) error
	
	// UpdateFluid updates an existing fluid.
	UpdateFluid(ctx context.Context, fluid *entities.Fluid) (*entities.Fluid, error)
	
	// DeleteFluid deletes a fluid by its ID.
	DeleteFluid(ctx context.Context, id string) error
}
