package repository

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

type TrajectoriesRepository interface {
	GetTrajectories(ctx context.Context, designID string) ([]*entities.Trajectory, error)
	GetTrajectoryByID(ctx context.Context, id string) (*entities.Trajectory, error)
	CreateTrajectory(ctx context.Context, designID string, trajectory *entities.Trajectory) error
	UpdateTrajectory(ctx context.Context, trajectory *entities.Trajectory) (*entities.Trajectory, error)
	DeleteTrajectory(ctx context.Context, id string) error
}