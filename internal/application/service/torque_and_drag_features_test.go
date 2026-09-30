package service

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

func fp(v float64) *float64 { return &v }

func section(kind string, top, bottom, od, id, weight float64) *entities.Section {
	return &entities.Section{Type: kind, BodyMD: bottom, BodyLength: bottom - top, BodyOD: od, BodyID: id, AvgJointLength: fp(9.14),
		Weight: fp(weight), MinYieldStrength: fp(105000)}
}

func stations(mds ...float64) []*entities.TrajectoryUnit {
	out := make([]*entities.TrajectoryUnit, len(mds))
	for i, md := range mds {
		out[i] = &entities.TrajectoryUnit{MD: md, TVD: md}
	}
	return out
}

func wellplanHoles() []*entities.Hole {
	return []*entities.Hole{{
		OpenHoleMDTop: 1000, OpenHoleMDBase: 2296, FrictionFactorOpenHole: 0.30,
		Caisings: []*entities.Caising{{MDTop: 0, MDBase: 1000, FrictionFactorCaising: 0.25}},
	}}
}

func TestBuildFeaturesMapsStationsToSectionBottoms(t *testing.T) {
	dp := section("Drill Pipe", 0, 1000, 127, 108.61, 32.62)
	dp.StabilizerOD = fp(152.4)
	secs := []*entities.Section{section("Bit", 1054, 1054.3, 215.9, 0, 100), dp, section("Heavy Weight", 1000, 1054, 127, 76.2, 73.13)}
	// Unsorted input, a duplicate, and stations below the bit.
	req, err := buildFeatures(stations(1020, 0, 500, 1000, 1000, 1054.3, 1500, 2000), secs, wellplanHoles())
	if err != nil {
		t.Fatal(err)
	}
	if want := []float64{0, 500, 1000, 1020, 1054.3}; !reflect.DeepEqual(req.MD, want) {
		t.Fatalf("stations: got %v want %v (sorted, deduplicated, cut at the bit)", req.MD, want)
	}
	if want := []float64{127, 127, 127, 127, 215.9}; !reflect.DeepEqual(req.BodyOD, want) {
		t.Fatalf("body OD by station: got %v want %v", req.BodyOD, want)
	}
	// Regression: the old code compared MD with the inner diameter (108.61), so every
	// station below 108 m fell back to the first or last section.
	if want := []float64{32.62, 32.62, 32.62, 73.13, 100}; !reflect.DeepEqual(req.Weight, want) {
		t.Fatalf("weight by station: got %v want %v", req.Weight, want)
	}
	if want := []float64{152.4, 152.4, 152.4, 0, 0}; !reflect.DeepEqual(req.StabilizerOD, want) {
		t.Fatalf("missing tool joint must be 0: got %v", req.StabilizerOD)
	}
	if want := []float64{0.25, 0.25, 0.25, 0.30, 0.30}; !reflect.DeepEqual(req.CoefficientOfFriction, want) {
		t.Fatalf("friction by hole section (casing wins at the shoe): got %v", req.CoefficientOfFriction)
	}
	for _, series := range [][]float64{req.Incl, req.Azim, req.SubSea, req.TVD, req.LocalNCoord, req.LocalECoord, req.GlobalNCoord,
		req.GlobalECoord, req.Dogleg, req.VerticalSection, req.BodyID, req.BodyAvgJointLength, req.StabilizerLength, req.StabilizerID, req.MinimumYieldStrength} {
		if len(series) != len(req.MD) {
			t.Fatal("every feature must have one value per station")
		}
	}
}

func TestBuildFeaturesReportsEveryProblemInsteadOfPanicking(t *testing.T) {
	bad := section("Drill Pipe", 0, 1000, 127, 108, 32)
	bad.Weight, bad.MinYieldStrength = nil, nil
	_, err := buildFeatures(stations(0, 500, 900), []*entities.Section{bad}, nil)
	var caseErr *CaseDataError
	if !errors.As(err, &caseErr) {
		t.Fatalf("want CaseDataError, got %v", err)
	}
	joined := strings.Join(caseErr.Problems, "\n")
	for _, want := range []string{"не указано — вес", "не указано — предел текучести", "нет коэффициента трения"} {
		if !strings.Contains(joined, want) {
			t.Errorf("problems should mention %q:\n%s", want, joined)
		}
	}
	for name, run := range map[string]func() error{
		"no string": func() error { _, err := buildFeatures(stations(0, 1), nil, nil); return err },
		"one station": func() error {
			_, err := buildFeatures(stations(5), []*entities.Section{section("DP", 0, 10, 1, 1, 1)}, wellplanHoles())
			return err
		},
		"gap": func() error {
			_, err := buildFeatures(stations(0, 5), []*entities.Section{section("A", 0, 10, 1, 1, 1), section("B", 15, 20, 1, 1, 1)}, wellplanHoles())
			return err
		},
		"overlap": func() error {
			_, err := buildFeatures(stations(0, 5), []*entities.Section{section("A", 0, 10, 1, 1, 1), section("B", 5, 20, 1, 1, 1)}, wellplanHoles())
			return err
		},
		"string too high": func() error {
			_, err := buildFeatures(stations(50, 60), []*entities.Section{section("DP", 0, 10, 1, 1, 1)}, wellplanHoles())
			return err
		},
	} {
		if err := run(); !errors.As(err, &caseErr) {
			t.Errorf("%s: want CaseDataError, got %v", name, err)
		}
	}
}

func TestSectionFrictionIsTheFallback(t *testing.T) {
	s := section("DP", 0, 1000, 127, 108, 32)
	s.FrictionCoefficient = fp(0.2)
	req, err := buildFeatures(stations(0, 500), []*entities.Section{s}, nil)
	if err != nil || req.CoefficientOfFriction[1] != 0.2 {
		t.Fatalf("got %v, %v", req, err)
	}
}

// Real WellPlan reports round lengths: the bit is printed as 0 m at the motor's depth and has
// no grade. It contains no station, so its missing yield must not block the prediction.
func TestZeroLengthBitAtMotorDepth(t *testing.T) {
	bit := section("Bit", 2296, 2296, 215.9, 0, 100)
	bit.MinYieldStrength = nil
	secs := []*entities.Section{bit, section("Drill Pipe", 0, 2286, 127, 108.61, 32.62), section("Mud Motor", 2286, 2296, 171.45, 76.2, 103.53)}
	req, err := buildFeatures(stations(0, 2290, 2296), secs, wellplanHoles())
	if err != nil {
		t.Fatal(err)
	}
	if want := []float64{127, 171.45, 171.45}; !reflect.DeepEqual(req.BodyOD, want) {
		t.Fatalf("station at 2296 m belongs to the motor: got %v", req.BodyOD)
	}
}
