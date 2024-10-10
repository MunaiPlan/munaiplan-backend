package postgres

import (
	"context"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"gorm.io/gorm"
)

type pressureDataProfilesRepository struct {
	db *gorm.DB
}

func NewPressureDataProfilesRepository(db *gorm.DB) *pressureDataProfilesRepository {
	return &pressureDataProfilesRepository{db: db}
}

func (r *pressureDataProfilesRepository) CreatePressureDataProfile(ctx context.Context, profile *entities.PressureDataProfile) error {
	gormProfile := toGormPressureDataProfile(profile)
	return r.db.WithContext(ctx).Create(gormProfile).Error
}

func (r *pressureDataProfilesRepository) GetPressureDataProfileByID(ctx context.Context, id string) (*entities.PressureDataProfile, error) {
	var profileModel models.PressureDataProfile
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&profileModel).Error
	if err != nil {
		return nil, err
	}
	return toDomainPressureDataProfile(&profileModel), nil
}

func (r *pressureDataProfilesRepository) GetPressureDataProfiles(ctx context.Context, caseID string) ([]*entities.PressureDataProfile, error) {
	var profileModels []models.PressureDataProfile
	err := r.db.WithContext(ctx).Where("case_id = ?", caseID).Find(&profileModels).Error
	if err != nil {
		return nil, err
	}

	profiles := make([]*entities.PressureDataProfile, len(profileModels))
	for i, profileModel := range profileModels {
		profiles[i] = toDomainPressureDataProfile(&profileModel)
	}
	return profiles, nil
}

func (r *pressureDataProfilesRepository) UpdatePressureDataProfile(ctx context.Context, profile *entities.PressureDataProfile) error {
	gormProfile := toGormPressureDataProfile(profile)
	return r.db.WithContext(ctx).Where("id = ?", profile.ID).Updates(gormProfile).Error
}

func (r *pressureDataProfilesRepository) DeletePressureDataProfile(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.PressureDataProfile{}).Error
}
