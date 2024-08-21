package service

import (
	"context"
	"fmt"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
)

type casesService struct {
	commonRepo repository.CommonRepository
	repo       repository.CasesRepository
}

func NewCasesService(repo repository.CasesRepository, commonRepo repository.CommonRepository) *casesService {
	return &casesService{
		repo:       repo,
		commonRepo: commonRepo,
	}
}

func (s *casesService) GetCases(ctx context.Context, input *requests.GetCasesRequest) ([]*entities.Case, error) {
	if err := s.commonRepo.CheckIfTrajectoryExists(ctx, input.TrajectoryID); err != nil {
		return nil, err
	}

	fmt.Println(input.TrajectoryID)

	return s.repo.GetCases(ctx, input.TrajectoryID)
}

func (s *casesService) GetCaseByID(ctx context.Context, input *requests.GetCaseByIDRequest) (*entities.Case, error) {
	return s.repo.GetCaseByID(ctx, input.ID)
}

func (s *casesService) CreateCase(ctx context.Context, input *requests.CreateCaseRequest) error {
	if err := s.commonRepo.CheckIfTrajectoryExists(ctx, input.TrajectoryID); err != nil {
		return err
	}

	caseEntity := &entities.Case{
		CaseName:        input.Body.CaseName,
		CaseDescription: input.Body.CaseDescription,
		DrillDepth:      input.Body.DrillDepth,
		PipeSize:        input.Body.PipeSize,
	}

	return s.repo.CreateCase(ctx, input.TrajectoryID, caseEntity)
}

func (s *casesService) UpdateCase(ctx context.Context, input *requests.UpdateCaseRequest) (*entities.Case, error) {
	caseEntity := &entities.Case{
		ID:              input.ID,
		CaseName:        input.Body.CaseName,
		CaseDescription: input.Body.CaseDescription,
		DrillDepth:      input.Body.DrillDepth,
		PipeSize:        input.Body.PipeSize,
	}

	return s.repo.UpdateCase(ctx, caseEntity)
}

func (s *casesService) DeleteCase(ctx context.Context, input *requests.DeleteCaseRequest) error {
	return s.repo.DeleteCase(ctx, input.ID)
}
