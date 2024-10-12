package requests

// CreateFractureGradientRequestBody represents the request body for creating a fracture gradient
type CreateFractureGradientRequestBody struct {
	TemperatureAtSurface float64 `json:"temperature_at_surface"`
	TemperatureAtWellTVD float64 `json:"temperature_at_well_tvd"`
	TemperatureGradient  float64 `json:"temperature_gradient"`
	WellTVD              float64 `json:"well_tvd"`
}

// CreateFractureGradientRequest represents the request for creating a fracture gradient
type CreateFractureGradientRequest struct {
	Body   CreateFractureGradientRequestBody
	CaseID string
}

// UpdateFractureGradientRequestBody represents the request body for updating a fracture gradient
type UpdateFractureGradientRequestBody struct {
	TemperatureAtSurface float64 `json:"temperature_at_surface"`
	TemperatureAtWellTVD float64 `json:"temperature_at_well_tvd"`
	TemperatureGradient  float64 `json:"temperature_gradient"`
	WellTVD              float64 `json:"well_tvd"`
}

// UpdateFractureGradientRequest represents the request for updating a fracture gradient
type UpdateFractureGradientRequest struct {
	ID   string
	Body UpdateFractureGradientRequestBody
}

// GetFractureGradientsRequest represents the request for getting fracture gradients
type GetFractureGradientsRequest struct {
	CaseID string
}

// GetFractureGradientByIDRequest represents the request for getting a fracture gradient by ID
type GetFractureGradientByIDRequest struct {
	ID string
}

// DeleteFractureGradientRequest represents the request for deleting a fracture gradient
type DeleteFractureGradientRequest struct {
	ID string
}
