package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type CasesRepository interface {
	CreateCase(ctx context.Context, trajectoryID string, caseEntity *entities.Case) error
	GetCaseByID(ctx context.Context, id string) (*entities.Case, error)
	GetCases(ctx context.Context, trajectoryID string) ([]*entities.Case, error)
	UpdateCase(ctx context.Context, caseEntity *entities.Case) (*entities.Case, error)
	DeleteCase(ctx context.Context, id string) error
}
