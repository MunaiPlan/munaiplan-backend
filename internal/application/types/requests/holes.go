package requests

// CreateCaisingRequestBody represents the request body for creating a caising associated with a hole.
type CreateCaisingRequestBody struct {
	MDTop                 float64  `json:"md_top"`
	MDBase                float64  `json:"md_base"`
	Length                float64  `json:"length"`
	ShoeMD                *float64 `json:"shoe_md,omitempty"`
	OD                    float64  `json:"od"`
	VD                    float64  `json:"vd"`
	DriftID               float64  `json:"drift_id"`
	EffectiveHoleDiameter float64  `json:"effective_hole_diameter"`
	Weight                float64  `json:"weight"`
	Grade                 string   `json:"grade"`
	MinYieldStrength      float64  `json:"min_yield_strength"`
	BurstRating           float64  `json:"burst_rating"`
	CollapseRating        float64  `json:"collapse_rating"`
	FrictionFactorCaising float64  `json:"friction_factor_caising"`
	LinearCapacityCaising float64  `json:"linear_capacity_caising"`
	DescriptionCaising    *string  `json:"description_caising,omitempty"`
	ManufacturerCaising   *string  `json:"manufacturer_caising,omitempty"`
	ModelCaising          *string  `json:"model_caising,omitempty"`
}

// UpdateCaisingRequestBody represents the request body for updating a caising associated with a hole.
type UpdateCaisingRequestBody struct {
	ID                    string  `json:"id"`
	MDTop                 float64  `json:"md_top"`
	MDBase                float64  `json:"md_base"`
	Length                float64  `json:"length"`
	ShoeMD                *float64 `json:"shoe_md,omitempty"`
	OD                    float64  `json:"od"`
	VD                    float64  `json:"vd"`
	DriftID               float64  `json:"drift_id"`
	EffectiveHoleDiameter float64  `json:"effective_hole_diameter"`
	Weight                float64  `json:"weight"`
	Grade                 string   `json:"grade"`
	MinYieldStrength      float64  `json:"min_yield_strength"`
	BurstRating           float64  `json:"burst_rating"`
	CollapseRating        float64  `json:"collapse_rating"`
	FrictionFactorCaising float64  `json:"friction_factor_caising"`
	LinearCapacityCaising float64  `json:"linear_capacity_caising"`
	DescriptionCaising    *string  `json:"description_caising,omitempty"`
	ManufacturerCaising   *string  `json:"manufacturer_caising,omitempty"`
	ModelCaising          *string  `json:"model_caising,omitempty"`
}

// CreateHoleRequestBody represents the request body for creating a hole.
type CreateHoleRequestBody struct {
	OpenHoleMDTop           float64                     `json:"open_hole_md_top"`
	OpenHoleMDBase          float64                     `json:"open_hole_md_base"`
	OpenHoleLength          float64                     `json:"open_hole_length"`
	OpenHoleVD              float64                     `json:"open_hole_vd"`
	EffectiveDiameter       float64                     `json:"effective_diameter"`
	FrictionFactorOpenHole  float64                     `json:"friction_factor_open_hole"`
	LinearCapacityOpenHole  float64                     `json:"linear_capacity_open_hole"`
	VolumeExcess            *float64                    `json:"volume_excess,omitempty"`
	DescriptionOpenHole     *string                     `json:"description_open_hole,omitempty"`
	TrippingInCasing        float64                     `json:"tripping_in_casing"`
	TrippingOutCasing       float64                     `json:"tripping_out_casing"`
	RotatingOnBottomCasing  float64                     `json:"rotating_on_bottom_casing"`
	SlideDrillingCasing     float64                     `json:"slide_drilling_casing"`
	BackReamingCasing       float64                     `json:"back_reaming_casing"`
	RotatingOffBottomCasing float64                     `json:"rotating_off_bottom_casing"`
	TrippingInOpenHole      float64                     `json:"tripping_in_open_hole"`
	TrippingOutOpenHole     float64                     `json:"tripping_out_open_hole"`
	RotatingOnBottomOpenHole float64                    `json:"rotating_on_bottom_open_hole"`
	SlideDrillingOpenHole   float64                     `json:"slide_drilling_open_hole"`
	BackReamingOpenHole     float64                     `json:"back_reaming_open_hole"`
	RotatingOffBottomOpenHole float64                   `json:"rotating_off_bottom_open_hole"`
	Caisings                []CreateCaisingRequestBody  `json:"caisings"`
}

// UpdateHoleRequestBody represents the request body for updating a hole.
type UpdateHoleRequestBody struct {
	OpenHoleMDTop           float64                     `json:"open_hole_md_top"`
	OpenHoleMDBase          float64                     `json:"open_hole_md_base"`
	OpenHoleLength          float64                     `json:"open_hole_length"`
	OpenHoleVD              float64                     `json:"open_hole_vd"`
	EffectiveDiameter       float64                     `json:"effective_diameter"`
	FrictionFactorOpenHole  float64                     `json:"friction_factor_open_hole"`
	LinearCapacityOpenHole  float64                     `json:"linear_capacity_open_hole"`
	VolumeExcess            *float64                    `json:"volume_excess,omitempty"`
	DescriptionOpenHole     *string                     `json:"description_open_hole,omitempty"`
	TrippingInCasing        float64                     `json:"tripping_in_casing"`
	TrippingOutCasing       float64                     `json:"tripping_out_casing"`
	RotatingOnBottomCasing  float64                     `json:"rotating_on_bottom_casing"`
	SlideDrillingCasing     float64                     `json:"slide_drilling_casing"`
	BackReamingCasing       float64                     `json:"back_reaming_casing"`
	RotatingOffBottomCasing float64                     `json:"rotating_off_bottom_casing"`
	TrippingInOpenHole      float64                     `json:"tripping_in_open_hole"`
	TrippingOutOpenHole     float64                     `json:"tripping_out_open_hole"`
	RotatingOnBottomOpenHole float64                    `json:"rotating_on_bottom_open_hole"`
	SlideDrillingOpenHole   float64                     `json:"slide_drilling_open_hole"`
	BackReamingOpenHole     float64                     `json:"back_reaming_open_hole"`
	RotatingOffBottomOpenHole float64                   `json:"rotating_off_bottom_open_hole"`
	Caisings                []UpdateCaisingRequestBody  `json:"caisings"`
}

// CreateHoleRequest represents the request for creating a hole.
type CreateHoleRequest struct {
		Body   CreateHoleRequestBody
	CaseID string
}

// UpdateHoleRequest represents the request for updating a hole.
type UpdateHoleRequest struct {
	ID    string
	Body  UpdateHoleRequestBody
}

// GetHolesRequest represents the request for getting holes by case ID.
type GetHolesRequest struct {
	CaseID string
}

// GetHoleByIDRequest represents the request for getting a hole by its ID.
type GetHoleByIDRequest struct {
	ID string
}

// DeleteHoleRequest represents the request for deleting a hole.
type DeleteHoleRequest struct {
	ID string
}
