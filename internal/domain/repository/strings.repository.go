package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type StringsRepository interface {
	// CreateString creates a new String along with its associated Sections.
	CreateString(ctx context.Context, caseID string, stringEntity *entities.String) error

	// GetStringByID retrieves a String by its ID, along with associated Sections.
	GetStringByID(ctx context.Context, id string) (*entities.String, error)

	// GetStrings retrieves all Strings associated with a case ID, along with associated Sections.
	GetStrings(ctx context.Context, caseID string) ([]*entities.String, error)

	// UpdateString updates an existing String and its associated Sections.
	UpdateString(ctx context.Context, stringEntity *entities.String) (*entities.String, error)

	// DeleteString deletes a String and its associated Sections from the database.
	DeleteString(ctx context.Context, id string) error
}
