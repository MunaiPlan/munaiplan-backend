package service

import (
	"context"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/responses"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
	client "github.com/munaiplan/munaiplan-backend/internal/infrastructure/prediction_client"
)

type torqueAndDragService struct {
	commonRepo repository.CommonRepository
	repo       repository.StringsRepository
	client     client.TorqueAndDragClient
}

func NewTorqueAndDragService(repo repository.StringsRepository, commonRepo repository.CommonRepository, client client.TorqueAndDragClient) *torqueAndDragService {
	return &torqueAndDragService{
		repo:       repo,
		commonRepo: commonRepo,
		client: client,
	}
}

func (s *torqueAndDragService) CalculateEffectiveTensionFromMLModel(ctx context.Context, caseID string) (*responses.EffectiveTensionFromMLModelResponse, error) {
	// Fetch trajectory data by case ID
	trajectory, err := s.commonRepo.GetTrajectoryByCaseID(ctx, caseID)
	if err != nil {
		return nil, err
	}

	// Fetch string data by case ID
	stringData, err := s.repo.GetStrings(ctx, caseID)
	if err != nil {
		return nil, err
	}

	// Map trajectory and string data to prepare for the client request
	mappedData := s.mapTrajectoryToStringSections(trajectory, stringData[0])

	// Call the external client with the mapped data
	response, err := s.client.CalculateEffectiveTension(mappedData)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/*
func (s *torqueAndDragService) WeightOnBitFromMlModel(ctx context.Context, input *requests.WeightOnBitFromMlModelRequest) (*entities.WeightOnBitFromMlModel, error) {
	return s.repo.WeightOnBitFromMlModel(ctx, input)
}

func (s *torqueAndDragService) HookLoadFromMlModel(ctx context.Context, input *requests.HookLoadFromMlModelRequest) (*entities.HookLoadFromMlModel, error) {
	return s.repo.HookLoadFromMlModel(ctx, input)
}

func (s *torqueAndDragService) SurfaceTorqueFromMlModel(ctx context.Context, input *requests.SurfaceTorqueFromMlModelRequest) (*entities.SurfaceTorqueFromMlModel, error) {
	return s.repo.SurfaceTorqueFromMlModel(ctx, input)
}
*/

// MapTrajectoryToStringSections maps data from String sections to Trajectory units based on MD depth.
func (s *torqueAndDragService) mapTrajectoryToStringSections(trajectory *entities.Trajectory, stringData *entities.String) requests.EffectiveTensionFromMLModelRequest {
	var result requests.EffectiveTensionFromMLModelRequest

	for _, unit := range trajectory.Units {
		// Append TrajectoryUnit attributes
		result.MD = append(result.MD, unit.MD)
		result.Incl = append(result.Incl, unit.Incl)
		result.Azim = append(result.Azim, unit.Azim)
		result.SubSea = append(result.SubSea, unit.SubSea)
		result.TVD = append(result.TVD, unit.TVD)
		result.LocalNCoord = append(result.LocalNCoord, unit.LocalNCoord)
		result.LocalECoord = append(result.LocalECoord, unit.LocalECoord)
		result.GlobalNCoord = append(result.GlobalNCoord, unit.GlobalNCoord)
		result.GlobalECoord = append(result.GlobalECoord, unit.GlobalECoord)
		result.Dogleg = append(result.Dogleg, unit.Dogleg)
		result.VerticalSection = append(result.VerticalSection, unit.VerticalSection)

		// Find the matching section based on MD range
		for _, section := range stringData.Sections {
			if unit.MD >= section.BodyMD-section.BodyLength && unit.MD <= section.BodyID {
				// Map Section data
				result.BodyOD = append(result.BodyOD, section.BodyOD)
				result.BodyID = append(result.BodyID, section.BodyID)
				result.BodyAvgJointLength = append(result.BodyAvgJointLength, *section.AvgJointLength)
				result.StabilizerLength = append(result.StabilizerLength, *section.StabilizerLength)
				result.StabilizerOD = append(result.StabilizerOD, *section.StabilizerOD)
				result.StabilizerID = append(result.StabilizerID, *section.StabilizerID)
				result.Weight = append(result.Weight, *section.Weight)
				result.CoefficientOfFriction = append(result.CoefficientOfFriction, *section.FrictionCoefficient)
				result.MinimumYieldStrength = append(result.MinimumYieldStrength, *section.MinYieldStrength)
				break
			}
		}
	}

	return result
}
