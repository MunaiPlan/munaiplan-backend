package entities

import "time"

// Hole entity
type Hole struct {
	ID                        string    `json:"id"`
	CaseID                    string    `json:"case_id"`
	CreatedAt                 time.Time `json:"created_at"`
	MDTop                     float64   `json:"md_top"`
	MDBase                    float64   `json:"md_base"`
	Length                    float64   `json:"length"`
	ShoeMD                    float64   `json:"shoe_md"`
	OD                        float64   `json:"od"`
	CaisingInternalDiameter   float64   `json:"caising_internal_diameter"`
	DriftInternalDiameter     float64   `json:"drift_internal_diameter"`
	EffectiveHoleDiameter     float64   `json:"effective_hole_diameter"`
	Weight                    float64   `json:"weight"`
	Grade                     string    `json:"grade"`
	MinYieldStrength          float64   `json:"min_yield_strength"`
	BurstRating               float64   `json:"burst_rating"`
	CollapseRating            float64   `json:"collapse_rating"`
	FrictionFactorCasing      float64   `json:"friction_factor_casing"`
	LinearCapacityCasing      float64   `json:"linear_capacity_casing"`
	DescriptionCasing         string    `json:"description_casing"`
	ManufacturerCasing        string    `json:"manufacturer_casing"`
	ModelCasing               string    `json:"model_casing"`
	OpenHoleMDTop             float64   `json:"open_hole_md_top"`
	OpenHoleMDBase            float64   `json:"open_hole_md_base"`
	OpenHoleLength            float64   `json:"open_hole_length"`
	OpenHoleInternalDiameter  float64   `json:"open_hole_internal_diameter"`
	EffectiveDiameter         float64   `json:"effective_diameter"`
	FrictionFactorOpenHole    float64   `json:"friction_factor_open_hole"`
	LinearCapacityOpenHole    float64   `json:"linear_capacity_open_hole"`
	VolumeExcess              float64   `json:"volume_excess"`
	DescriptionOpenHole       string    `json:"description_open_hole"`
	TrippingInCasing          float64   `json:"tripping_in_casing"`
	TrippingOutCasing         float64   `json:"tripping_out_casing"`
	RotatingOnBottomCasing    float64   `json:"rotating_on_bottom_casing"`
	SlideDrillingCasing       float64   `json:"slide_drilling_casing"`
	BackReamingCasing         float64   `json:"back_reaming_casing"`
	RotatingOffBottomCasing   float64   `json:"rotating_off_bottom_casing"`
	TrippingInOpenHole        float64   `json:"tripping_in_open_hole"`
	TrippingOutOpenHole       float64   `json:"tripping_out_open_hole"`
	RotatingOnBottomOpenHole  float64   `json:"rotating_on_bottom_open_hole"`
	SlideDrillingOpenHole     float64   `json:"slide_drilling_open_hole"`
	BackReamingOpenHole       float64   `json:"back_reaming_open_hole"`
	RotatingOffBottomOpenHole float64   `json:"rotating_off_bottom_open_hole"`
}
