package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type DatumsRepository interface {
	CreateDatum(ctx context.Context, caseID string, datum *entities.Datum) error
	GetDatumByID(ctx context.Context, id string) (*entities.Datum, error)
	GetDatumsByCaseID(ctx context.Context, caseID string) ([]*entities.Datum, error)
	UpdateDatum(ctx context.Context, datum *entities.Datum) (*entities.Datum, error)
	DeleteDatum(ctx context.Context, id string) error
}
