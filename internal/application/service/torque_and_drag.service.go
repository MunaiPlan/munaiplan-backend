package service

import (
	"context"
	"sort"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/application/types/responses"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
	client "github.com/munaiplan/munaiplan-backend/internal/infrastructure/prediction_client"
)

type torqueAndDragService struct {
	commonRepo repository.CommonRepository
	strings    repository.StringsRepository
	holes      repository.HolesRepository
	imports    repository.ImportsRepository
	client     client.TorqueAndDragClient
}

func NewTorqueAndDragService(strings repository.StringsRepository, holes repository.HolesRepository, imports repository.ImportsRepository, commonRepo repository.CommonRepository, client client.TorqueAndDragClient) *torqueAndDragService {
	return &torqueAndDragService{strings: strings, holes: holes, imports: imports, commonRepo: commonRepo, client: client}
}

func (s *torqueAndDragService) CalculateEffectiveTensionFromMLModel(ctx context.Context, caseID string) (*responses.EffectiveTensionFromMLModelResponse, error) {
	var out responses.EffectiveTensionFromMLModelResponse
	depth, err := s.predict(ctx, caseID, client.EffectiveTension, &out)
	out.Depth = depth
	return &out, err
}

func (s *torqueAndDragService) CalculateWeightOnBitFromMlModel(ctx context.Context, caseID string) (*responses.WeightOnBitFromMLModelResponse, error) {
	var out responses.WeightOnBitFromMLModelResponse
	depth, err := s.predict(ctx, caseID, client.HookLoad, &out)
	out.Depth = depth
	return &out, err
}

func (s *torqueAndDragService) CalculateSurfaceTorqueFromMlModel(ctx context.Context, caseID string) (*responses.MomentFromMLModelResponse, error) {
	var out responses.MomentFromMLModelResponse
	depth, err := s.predict(ctx, caseID, client.Torque, &out)
	out.Depth = depth
	return &out, err
}

func (s *torqueAndDragService) CalculateMinWeightFromMLModel(ctx context.Context, caseID string) (*responses.MinWeightFromMLModelResponse, error) {
	var out responses.MinWeightFromMLModelResponse
	depth, err := s.predict(ctx, caseID, client.MinWeight, &out)
	out.Depth = depth
	return &out, err
}

// predict assembles the case's features, calls the model and returns the station depths.
func (s *torqueAndDragService) predict(ctx context.Context, caseID string, family client.Family, out any) ([]float64, error) {
	features, err := s.features(ctx, caseID)
	if err != nil {
		return nil, err
	}
	if err := s.client.Predict(ctx, family, *features, out); err != nil {
		return nil, err
	}
	return features.MD, nil
}

// features loads the trajectory, work string and hole sections of a case.
// With several strings, the oldest is used so results are deterministic.
func (s *torqueAndDragService) features(ctx context.Context, caseID string) (*requests.TorqueAndDragFromMLModelRequest, error) {
	trajectory, err := s.commonRepo.GetTrajectoryByCaseID(ctx, caseID)
	if err != nil {
		return nil, err
	}
	strs, err := s.strings.GetStrings(ctx, caseID)
	if err != nil {
		return nil, err
	}
	if len(strs) == 0 {
		return nil, &CaseDataError{[]string{"у кейса нет рабочей колонны"}}
	}
	sort.SliceStable(strs, func(i, j int) bool { return strs[i].CreatedAt.Before(strs[j].CreatedAt) })
	holes, err := s.holes.GetHoles(ctx, caseID)
	if err != nil {
		return nil, err
	}
	return buildFeatures(trajectory.Units, strs[0].Sections, holes)
}

// ModelReady reports whether predictions can currently be served.
func (s *torqueAndDragService) ModelReady(ctx context.Context) error {
	return s.client.Ready(ctx)
}
