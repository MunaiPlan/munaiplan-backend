package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type trajectoriesService struct {
	commonRepo repository.CommonRepository
	repo       repository.TrajectoriesRepository
}

func NewTrajectoriesService(repo repository.TrajectoriesRepository, commonRepo repository.CommonRepository) *trajectoriesService {
	return &trajectoriesService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *trajectoriesService) GetTrajectories(ctx context.Context, input *requests.GetTrajectoriesRequest) ([]*entities.Trajectory, error) {
	if err := s.commonRepo.CheckIfDesignExists(ctx, input.DesignID); err != nil {
		return nil, err
	}

	return s.repo.GetTrajectories(ctx, input.DesignID)
}

func (s *trajectoriesService) GetTrajectoryByID(ctx context.Context, input *requests.GetTrajectoryByIDRequest) (*entities.Trajectory, error) {
	return s.repo.GetTrajectoryByID(ctx, input.ID)
}

func (s *trajectoriesService) CreateTrajectory(ctx context.Context, input *requests.CreateTrajectoryRequest) error {
	if err := s.commonRepo.CheckIfDesignExists(ctx, input.DesignID); err != nil {
		return err
	}

	trajectory := s.CreateTrajectoryRequestToEntity(&input.Body)
	return s.repo.CreateTrajectory(ctx, input.DesignID, trajectory)
}

func (s *trajectoriesService) UpdateTrajectory(ctx context.Context, input *requests.UpdateTrajectoryRequest) (*entities.Trajectory, error) {
	trajectory := s.UpdateTrajectoryRequestToEntity(&input.Body)
	trajectory.ID = input.ID
	return s.repo.UpdateTrajectory(ctx, trajectory)
}

func (s *trajectoriesService) DeleteTrajectory(ctx context.Context, input *requests.DeleteTrajectoryRequest) error {
	return s.repo.DeleteTrajectory(ctx, input.ID)
}

func (s *trajectoriesService) CreateTrajectoryRequestToEntity(input *requests.CreateTrajectoryRequestBody) *entities.Trajectory {
	headers := make([]*entities.TrajectoryHeader, len(input.Headers))
	for i, header := range input.Headers {
		headers[i] = &entities.TrajectoryHeader{
			Customer:         header.Customer,
			Project:          header.Project,
			ProfileType:      header.ProfileType,
			Field:            header.Field,
			YourRef:          header.YourRef,
			Structure:        header.Structure,
			JobNumber:        header.JobNumber,
			Wellhead:         header.Wellhead,
			KellyBushingElev: header.KellyBushingElev,
			Profile:          header.Profile,
		}
	}

	units := make([]*entities.TrajectoryUnit, len(input.Units))
	for i, unit := range input.Units {
		units[i] = &entities.TrajectoryUnit{
			MD:              unit.MD,
			Incl:            unit.Incl,
			Azim:            unit.Azim,
			SubSea:          unit.SubSea,
			TVD:             unit.TVD,
			LocalNCoord:     unit.LocalNCoord,
			LocalECoord:     unit.LocalECoord,
			GlobalNCoord:    unit.GlobalNCoord,
			GlobalECoord:    unit.GlobalECoord,
			Dogleg:          unit.Dogleg,
			VerticalSection: unit.VerticalSection,
		}
	}

	trajectory := &entities.Trajectory{
		Name:        input.Name,
		Description: input.Description,
		Headers:     headers,
		Units:       units,
	}

	return trajectory
}

func (s *trajectoriesService) UpdateTrajectoryRequestToEntity(input *requests.UpdateTrajectoryRequestBody) *entities.Trajectory {
	headers := make([]*entities.TrajectoryHeader, len(input.Headers))
	for i, header := range input.Headers {
		headers[i] = &entities.TrajectoryHeader{
			ID:               header.ID,
			Customer:         header.Customer,
			Project:          header.Project,
			ProfileType:      header.ProfileType,
			Field:            header.Field,
			YourRef:          header.YourRef,
			Structure:        header.Structure,
			JobNumber:        header.JobNumber,
			Wellhead:         header.Wellhead,
			KellyBushingElev: header.KellyBushingElev,
			Profile:          header.Profile,
		}
	}

	units := make([]*entities.TrajectoryUnit, len(input.Units))
	for i, unit := range input.Units {
		units[i] = &entities.TrajectoryUnit{
			MD:              unit.MD,
			Incl:            unit.Incl,
			Azim:            unit.Azim,
			SubSea:          unit.SubSea,
			TVD:             unit.TVD,
			LocalNCoord:     unit.LocalNCoord,
			LocalECoord:     unit.LocalECoord,
			GlobalNCoord:    unit.GlobalNCoord,
			GlobalECoord:    unit.GlobalECoord,
			Dogleg:          unit.Dogleg,
			VerticalSection: unit.VerticalSection,
		}
		if unit.ID != nil {
			units[i].ID = *unit.ID
		} else {
			units[i].ID = ""
		}
	}

	trajectory := &entities.Trajectory{
		Name:        input.Name,
		Description: input.Description,
		Headers:     headers,
		Units:       units,
	}

	return trajectory
}
