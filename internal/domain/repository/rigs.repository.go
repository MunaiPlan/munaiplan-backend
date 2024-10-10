package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type RigsRepository interface {
	GetRigs(ctx context.Context, caseID string) ([]*entities.Rig, error)
	GetRigByID(ctx context.Context, id string) (*entities.Rig, error)
	CreateRig(ctx context.Context, caseID string, rig *entities.Rig) error
	UpdateRig(ctx context.Context, rig *entities.Rig) (*entities.Rig, error)
	DeleteRig(ctx context.Context, id string) error
}
