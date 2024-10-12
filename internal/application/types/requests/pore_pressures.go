package requests

// CreatePorePressureRequest represents the request body for creating a pore pressure record.
type CreatePorePressureRequest struct {
	CaseID string
	Body   CreatePorePressureRequestBody
}

type CreatePorePressureRequestBody struct {
	TVD      float64 `json:"tvd" binding:"required"`
	Pressure float64 `json:"pressure" binding:"required"`
	EMW      float64 `json:"emw" binding:"required"`
}

// UpdatePorePressureRequest represents the request body for updating a pore pressure record.
type UpdatePorePressureRequest struct {
	ID   string
	Body UpdatePorePressureRequestBody
}

type UpdatePorePressureRequestBody struct {
	TVD      float64 `json:"tvd" binding:"required"`
	Pressure float64 `json:"pressure" binding:"required"`
	EMW      float64 `json:"emw" binding:"required"`
}

// GetPorePressuresRequest represents the request for retrieving pore pressures associated with a case.
type GetPorePressuresRequest struct {
	CaseID string
}

// GetPorePressureByIDRequest represents the request for retrieving a pore pressure by ID.
type GetPorePressureByIDRequest struct {
	ID string
}

// DeletePorePressureRequest represents the request for deleting a pore pressure by ID.
type DeletePorePressureRequest struct {
	ID string
}
