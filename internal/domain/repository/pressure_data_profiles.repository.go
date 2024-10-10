package repository

import (
	"context"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type PressureDataProfilesRepository interface {
	CreatePressureDataProfile(ctx context.Context, profile *entities.PressureDataProfile) error
	GetPressureDataProfileByID(ctx context.Context, id string) (*entities.PressureDataProfile, error)
	GetPressureDataProfiles(ctx context.Context, caseID string) ([]*entities.PressureDataProfile, error)
	UpdatePressureDataProfile(ctx context.Context, profile *entities.PressureDataProfile) error
	DeletePressureDataProfile(ctx context.Context, id string) error
}
