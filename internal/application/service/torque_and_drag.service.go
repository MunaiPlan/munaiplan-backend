package service

import (
	"context"
	"sort"

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
	mappedData, err := s.getMappedRequestForCase(ctx, caseID)
	if err != nil {
		return nil, err
	}

	// Call the external client with the mapped data
	response, err := s.client.CalculateEffectiveTension(*mappedData)
	if err != nil {
		return nil, err
	}

	response.Depth = mappedData.MD

	return response, nil
}

// func (s *torqueAndDragService) WeightOnBitFromMlModel(ctx context.Context, caseID string) (*responses.WeightOnBitFromMlModel, error) {
// 	mappedData, err := s.getMappedRequestForCase(ctx, caseID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Call the external client with the mapped data
// 	response, err := s.client.CalculateEffectiveTension(*mappedData)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response.Depth = mappedData.MD

// 	return response, nil
// }

func (s *torqueAndDragService) getMappedRequestForCase(ctx context.Context, caseID string) (*requests.EffectiveTensionFromMLModelRequest, error) {
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

	return &mappedData, nil
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
	sort.Slice(stringData.Sections, func(i, j int) bool {
		return stringData.Sections[i].BodyMD < stringData.Sections[j].BodyMD
	})

	for _, unit := range trajectory.Units {
		var found bool = false
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
				found = true
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

		if !found {
			if unit.MD < stringData.Sections[0].BodyMD {
				// Map Section data
				result.BodyOD = append(result.BodyOD, stringData.Sections[0].BodyOD)
				result.BodyID = append(result.BodyID, stringData.Sections[0].BodyID)
				result.BodyAvgJointLength = append(result.BodyAvgJointLength, *stringData.Sections[0].AvgJointLength)
				result.StabilizerLength = append(result.StabilizerLength, *stringData.Sections[0].StabilizerLength)
				result.StabilizerOD = append(result.StabilizerOD, *stringData.Sections[0].StabilizerOD)
				result.StabilizerID = append(result.StabilizerID, *stringData.Sections[0].StabilizerID)
				result.Weight = append(result.Weight, *stringData.Sections[0].Weight)
				result.CoefficientOfFriction = append(result.CoefficientOfFriction, *stringData.Sections[0].FrictionCoefficient)
				result.MinimumYieldStrength = append(result.MinimumYieldStrength, *stringData.Sections[0].MinYieldStrength)
			} else {
				// Map Section data
				result.BodyOD = append(result.BodyOD, stringData.Sections[len(stringData.Sections)-1].BodyOD)
				result.BodyID = append(result.BodyID, stringData.Sections[len(stringData.Sections)-1].BodyID)
				result.BodyAvgJointLength = append(result.BodyAvgJointLength, *stringData.Sections[len(stringData.Sections)-1].AvgJointLength)
				result.StabilizerLength = append(result.StabilizerLength, *stringData.Sections[len(stringData.Sections)-1].StabilizerLength)
				result.StabilizerOD = append(result.StabilizerOD, *stringData.Sections[len(stringData.Sections)-1].StabilizerOD)
				result.StabilizerID = append(result.StabilizerID, *stringData.Sections[len(stringData.Sections)-1].StabilizerID)
				result.Weight = append(result.Weight, *stringData.Sections[len(stringData.Sections)-1].Weight)
				result.CoefficientOfFriction = append(result.CoefficientOfFriction, *stringData.Sections[len(stringData.Sections)-1].FrictionCoefficient)
				result.MinimumYieldStrength = append(result.MinimumYieldStrength, *stringData.Sections[len(stringData.Sections)-1].MinYieldStrength)
			}
		}
	}

	return result
}
