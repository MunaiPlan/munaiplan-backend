package service

import (
	"testing"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/responses"
	"github.com/munaiplan/munaiplan-backend/internal/importers/wellplan"
)

func TestCompareReadsHookLoadAtBitAndTorqueAtSurface(t *testing.T) {
	report := &wellplan.Report{TorqueDrag: &wellplan.TorqueDragResults{Loads: []wellplan.LoadCondition{
		{Operation: wellplan.OpTrippingOut, HookLoad: fp(80), SurfaceTorque: fp(0)},
		{Operation: wellplan.OpRotatingOnBottom, HookLoad: fp(68), SurfaceTorque: fp(10)},
		{Operation: wellplan.OpSlideDrilling, HookLoad: fp(63), SurfaceTorque: fp(0)},
		{Operation: wellplan.OpRotatingOffBottom, HookLoad: fp(74), SurfaceTorque: fp(4.7)},
	}}}
	hook := &responses.WeightOnBitFromMLModelResponse{PullUp: []float64{30, 60, 72}, RotaryDrilling: []float64{20, 50, 70}, DrillingGZD: []float64{1, 2, 3}}
	torque := &responses.MomentFromMLModelResponse{PullUp: []float64{0.1, 0, 0}, RotaryDrilling: []float64{9, 5, 1}}
	got := compare("c", []float64{0, 1000, 2296}, report, hook, torque)
	if got.BitDepth != 2296 || len(got.Rows) != 8 {
		t.Fatalf("rows: %+v", got.Rows)
	}
	pooh := got.Rows[0]
	if *pooh.Model != 72 || *pooh.Difference != -8 || *pooh.RelativeDifference != -0.1 {
		t.Fatalf("hook load must be read at the bit: %+v", pooh)
	}
	rotTorque := got.Rows[3]
	if rotTorque.Metric != "surface_torque" || *rotTorque.Model != 9 || *rotTorque.Difference != -1 {
		t.Fatalf("torque must be read at the surface: %+v", rotTorque)
	}
	if got.Rows[1].RelativeDifference != nil {
		t.Fatal("no relative difference against a zero reference")
	}
	if got.Rows[5].Model != nil || got.Rows[6].Model != nil {
		t.Fatal("operations without a model series must have no model value")
	}
}
