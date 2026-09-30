package service

import (
	"context"
	"math"
	"sync"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/responses"
	"github.com/munaiplan/munaiplan-backend/internal/importers/wellplan"
	client "github.com/munaiplan/munaiplan-backend/internal/infrastructure/prediction_client"
	"gorm.io/gorm"
)

var comparisonNotes = []string{
	"Вес на крюке: значение из отчёта на поверхности для анализируемой глубины долота сравнивается с кривой модели в самой глубокой точке колонны.",
	"Момент на поверхности: значение из отчёта на роторе сравнивается с кривой момента модели в самой верхней точке.",
	"Эталон — это расчёт исходной инженерной программы, а не замеры на буровой. Совпадение — проверка согласованности, а не валидация.",
	"Единицы выхода модели не документированы; значения модели показаны как есть.",
}

// modelSeries selects the model output that corresponds to a WellPlan operation.
type modelSeries struct {
	hookLoad func(*responses.WeightOnBitFromMLModelResponse) []float64
	torque   func(*responses.MomentFromMLModelResponse) []float64
}

var seriesByOperation = map[string]modelSeries{
	wellplan.OpTrippingIn: {
		func(r *responses.WeightOnBitFromMLModelResponse) []float64 { return r.RunIn },
		func(r *responses.MomentFromMLModelResponse) []float64 { return r.RunIn }},
	wellplan.OpTrippingOut: {
		func(r *responses.WeightOnBitFromMLModelResponse) []float64 { return r.PullUp },
		func(r *responses.MomentFromMLModelResponse) []float64 { return r.PullUp }},
	wellplan.OpRotatingOnBottom: {
		func(r *responses.WeightOnBitFromMLModelResponse) []float64 { return r.RotaryDrilling },
		func(r *responses.MomentFromMLModelResponse) []float64 { return r.RotaryDrilling }},
	wellplan.OpSlideDrilling: {
		func(r *responses.WeightOnBitFromMLModelResponse) []float64 { return r.DrillingGZD },
		nil}, // the torque model has no slide-drilling series
}

// CompareWithReference runs the hook-load and torque models for an imported case and lines
// them up with WellPlan's load summary stored at import time.
func (s *torqueAndDragService) CompareWithReference(ctx context.Context, organizationID, caseID string) (*responses.ReferenceComparison, error) {
	report, err := s.imports.GetReport(ctx, organizationID, caseID)
	if err != nil {
		return nil, err
	}
	if report == nil || report.TorqueDrag == nil || len(report.TorqueDrag.Loads) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	features, err := s.features(ctx, caseID)
	if err != nil {
		return nil, err
	}
	// The two model calls are independent; run them together to halve the latency.
	var (
		hook             responses.WeightOnBitFromMLModelResponse
		torque           responses.MomentFromMLModelResponse
		hookErr, torqErr error
		wg               sync.WaitGroup
	)
	wg.Add(2)
	go func() { defer wg.Done(); hookErr = s.client.Predict(ctx, client.HookLoad, *features, &hook) }()
	go func() { defer wg.Done(); torqErr = s.client.Predict(ctx, client.Torque, *features, &torque) }()
	wg.Wait()
	if hookErr != nil {
		return nil, hookErr
	}
	if torqErr != nil {
		return nil, torqErr
	}
	return compare(caseID, features.MD, report, &hook, &torque), nil
}

func compare(caseID string, depth []float64, report *wellplan.Report, hook *responses.WeightOnBitFromMLModelResponse, torque *responses.MomentFromMLModelResponse) *responses.ReferenceComparison {
	out := &responses.ReferenceComparison{CaseID: caseID, BitDepth: depth[len(depth)-1], Notes: comparisonNotes, Warnings: report.Warnings}
	for _, load := range report.TorqueDrag.Loads {
		series, known := seriesByOperation[load.Operation]
		var modelHook, modelTorque *float64
		if known && series.hookLoad != nil {
			modelHook = last(series.hookLoad(hook))
		}
		if known && series.torque != nil {
			modelTorque = first(series.torque(torque))
		}
		out.Rows = append(out.Rows,
			row(load.Operation, "hook_load", "t", load.HookLoad, modelHook, "hook-load curve at the deepest station"),
			row(load.Operation, "surface_torque", "kN-m", load.SurfaceTorque, modelTorque, "torque curve at the shallowest station"))
	}
	return out
}

func row(operation, metric, unit string, reference, model *float64, point string) responses.ComparisonRow {
	r := responses.ComparisonRow{Operation: operation, Metric: metric, Unit: unit, WellPlan: reference, Model: model, ModelPoint: point}
	if reference != nil && model != nil {
		diff := *model - *reference
		r.Difference = &diff
		if math.Abs(*reference) > 1e-9 {
			rel := diff / math.Abs(*reference)
			r.RelativeDifference = &rel
		}
	}
	return r
}

func first(values []float64) *float64 {
	if len(values) == 0 {
		return nil
	}
	return &values[0]
}

func last(values []float64) *float64 {
	if len(values) == 0 {
		return nil
	}
	return &values[len(values)-1]
}
