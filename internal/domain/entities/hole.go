package entities

import "time"

// Hole entity
type Hole struct {
	ID        string            `json:"id"`
	CaseID    string            `json:"case_id"`
	CreatedAt time.Time         `json:"created_at"`
	Casings   []*HoleCasing     `json:"casings"`
	OpenHoles []*OpenHole       `json:"open_holes"`
	Factors   []*FrictionFactor `json:"friction_factors"`
}

// HoleCasing entity
type HoleCasing struct {
	ID                    string    `json:"id"`
	HoleID                string    `json:"hole_id"`
	MDTop                 float64   `json:"md_top"`
	MDBase                float64   `json:"md_base"`
	Length                float64   `json:"length"`
	ShoeMD                float64   `json:"shoe_md"`
	OD                    float64   `json:"od"`
	IDValue               float64   `json:"id_value"` // To avoid collision with ID field
	DriftID               float64   `json:"drift_id"`
	EffectiveHoleDiameter float64   `json:"effective_hole_diameter"`
	Weight                float64   `json:"weight"`
	Grade                 string    `json:"grade"`
	MinYieldStrength      float64   `json:"min_yield_strength"`
	BurstRating           float64   `json:"burst_rating"`
	CollapseRating        float64   `json:"collapse_rating"`
	FrictionFactor        float64   `json:"friction_factor"`
	LinearCapacity        float64   `json:"linear_capacity"`
	Description           string    `json:"description"`
	Manufacturer          string    `json:"manufacturer"`
	Model                 string    `json:"model"`
	CreatedAt             time.Time `json:"created_at"`
}

// OpenHole entity
type OpenHole struct {
	ID                string    `json:"id"`
	HoleID            string    `json:"hole_id"`
	MDTop             float64   `json:"md_top"`
	MDBase            float64   `json:"md_base"`
	Length            float64   `json:"length"`
	IDValue           float64   `json:"id_value"` // To avoid collision with ID field
	EffectiveDiameter float64   `json:"effective_diameter"`
	FrictionFactor    float64   `json:"friction_factor"`
	LinearCapacity    float64   `json:"linear_capacity"`
	VolumeExcess      float64   `json:"volume_excess"`
	Description       string    `json:"description"`
	CreatedAt         time.Time `json:"created_at"`
}

// FrictionFactor entity
type FrictionFactor struct {
	ID                        string    `json:"id"`
	HoleID                    string    `json:"hole_id"`
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
	CreatedAt                 time.Time `json:"created_at"`
}
