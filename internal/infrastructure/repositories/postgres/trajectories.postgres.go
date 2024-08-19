package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"gorm.io/gorm"
)

type trajectoriesRepository struct {
	db *gorm.DB
}

func NewTrajectoriesRepository(db *gorm.DB) *trajectoriesRepository {
	return &trajectoriesRepository{db: db}
}

func (r *trajectoriesRepository) CreateTrajectory(ctx context.Context, designID string, trajectory *entities.Trajectory) error {
	gormTrajectory := r.toGormTrajectory(trajectory)
	designId, err := uuid.Parse(designID)
	if err != nil {
		return err
	}
	gormTrajectory.DesignID = designId

	result := r.db.WithContext(ctx).Create(gormTrajectory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *trajectoriesRepository) GetTrajectoryByID(ctx context.Context, id string) (*entities.Trajectory, error) {
	var trajectory models.Trajectory
	var res entities.Trajectory
	result := r.db.WithContext(ctx).Preload("Headers").Preload("Units").Where("id = ?", id).First(&trajectory)
	if result.Error != nil {
		return nil, result.Error
	}

	res = r.toDomainTrajectory(&trajectory)
	return &res, nil
}

func (r *trajectoriesRepository) GetTrajectories(ctx context.Context, designID string) ([]*entities.Trajectory, error) {
	var trajectories []*models.Trajectory
	var res []*entities.Trajectory
	result := r.db.WithContext(ctx).Preload("Headers").Preload("Units").Where("design_id = ?", designID).Find(&trajectories)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, trajectory := range trajectories {
		temp := r.toDomainTrajectory(trajectory)
		res = append(res, &temp)
	}
	return res, nil
}

func (r *trajectoriesRepository) UpdateTrajectory(ctx context.Context, trajectory *entities.Trajectory) (*entities.Trajectory, error) {
	gormTrajectory := r.toGormTrajectory(trajectory)
	var oldTrajectory, updatedTrajectory models.Trajectory

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Preload("Headers").Preload("Units").Where("id = ?", trajectory.ID).First(&oldTrajectory).Error; err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Model(&oldTrajectory).Updates(gormTrajectory).Error; err != nil {
			return err
		}
		if err := tx.Model(&oldTrajectory).Association("Headers").Replace(gormTrajectory.Headers); err != nil {
			return err
		}
		if err := tx.Model(&oldTrajectory).Association("Units").Replace(gormTrajectory.Units); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Preload("Headers").Preload("Units").Where("id = ?", trajectory.ID).First(&updatedTrajectory).Error; err != nil {
		return nil, err
	}

	res := r.toDomainTrajectory(&updatedTrajectory)
	return &res, nil
}

func (r *trajectoriesRepository) DeleteTrajectory(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Trajectory{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// toDomainTrajectory maps the GORM Trajectory model to the domain Trajectory entity.
func (r *trajectoriesRepository) toDomainTrajectory(trajectoryModel *models.Trajectory) entities.Trajectory {
	headers := make([]*entities.TrajectoryHeader, len(trajectoryModel.Headers))
	for i, h := range trajectoryModel.Headers {
		headers[i] = &entities.TrajectoryHeader{
			ID:               h.ID.String(),
			Customer:         h.Customer,
			Project:          h.Project,
			ProfileType:      h.ProfileType,
			Field:            h.Field,
			YourRef:          h.YourRef,
			Structure:        h.Structure,
			JobNumber:        h.JobNumber,
			Wellhead:         h.Wellhead,
			KellyBushingElev: h.KellyBushingElev,
			Profile:          h.Profile,
			CreatedAt:        h.CreatedAt,
		}
	}

	units := make([]*entities.TrajectoryUnit, len(trajectoryModel.Units))
	for i, u := range trajectoryModel.Units {
		units[i] = &entities.TrajectoryUnit{
			ID:              u.ID.String(),
			MD:              u.MD,
			Incl:            u.Incl,
			Azim:            u.Azim,
			SubSea:          u.SubSea,
			TVD:             u.TVD,
			LocalNCoord:     u.LocalNCoord,
			LocalECoord:     u.LocalECoord,
			GlobalNCoord:    u.GlobalNCoord,
			GlobalECoord:    u.GlobalECoord,
			Dogleg:          u.Dogleg,
			VerticalSection: u.VerticalSection,
			CreatedAt:       u.CreatedAt,
		}
	}

	return entities.Trajectory{
		ID:          trajectoryModel.ID.String(),
		Name:        trajectoryModel.Name,
		Description: trajectoryModel.Description,
		Headers:     headers,
		Units:       units,
		CreatedAt:   trajectoryModel.CreatedAt,
	}
}

// toGormTrajectory maps the domain Trajectory entity to the GORM Trajectory model.
func (r *trajectoriesRepository) toGormTrajectory(trajectory *entities.Trajectory) *models.Trajectory {
	headers := make([]models.TrajectoryHeader, len(trajectory.Headers))
	for i, h := range trajectory.Headers {
		headerID, _ := uuid.Parse(h.ID)
		headers[i] = models.TrajectoryHeader{
			ID:               headerID,
			Customer:         h.Customer,
			Project:          h.Project,
			ProfileType:      h.ProfileType,
			Field:            h.Field,
			YourRef:          h.YourRef,
			Structure:        h.Structure,
			JobNumber:        h.JobNumber,
			Wellhead:         h.Wellhead,
			KellyBushingElev: h.KellyBushingElev,
			Profile:          h.Profile,
			CreatedAt:        h.CreatedAt,
		}
	}

	units := make([]models.TrajectoryUnit, len(trajectory.Units))
	for i, u := range trajectory.Units {
		unitID, _ := uuid.Parse(u.ID)
		units[i] = models.TrajectoryUnit{
			ID:              unitID,
			MD:              u.MD,
			Incl:            u.Incl,
			Azim:            u.Azim,
			SubSea:          u.SubSea,
			TVD:             u.TVD,
			LocalNCoord:     u.LocalNCoord,
			LocalECoord:     u.LocalECoord,
			GlobalNCoord:    u.GlobalNCoord,
			GlobalECoord:    u.GlobalECoord,
			Dogleg:          u.Dogleg,
			VerticalSection: u.VerticalSection,
			CreatedAt:       u.CreatedAt,
		}
	}

	return &models.Trajectory{
		Name:        trajectory.Name,
		Description: trajectory.Description,
		Headers:     headers,
		Units:       units,
	}
}
