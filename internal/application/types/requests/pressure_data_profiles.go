package requests

// CreatePressureDataProfileRequest represents the request to create a new pressure data profile.
type CreatePressureDataProfileRequest struct {
	CaseID string
	Body   CreatePressureDataProfileRequestBody
}

type CreatePressureDataProfileRequestBody struct {
	TVD      float64 `json:"tvd"`
	Pressure float64 `json:"pressure"`
	EMW      float64 `json:"emw"`
}

// UpdatePressureDataProfileRequest represents the request to update an existing pressure data profile.
type UpdatePressureDataProfileRequest struct {
	ID   string
	Body UpdatePressureDataProfileRequestBody
}

type UpdatePressureDataProfileRequestBody struct {
	TVD      float64 `json:"tvd"`
	Pressure float64 `json:"pressure"`
	EMW      float64 `json:"emw"`
}

// GetPressureDataProfilesRequest represents the request to fetch pressure data profiles for a case.
type GetPressureDataProfilesRequest struct {
	CaseID string
}

// GetPressureDataProfileByIDRequest represents the request to fetch a pressure data profile by ID.
type GetPressureDataProfileByIDRequest struct {
	ID string
}

// DeletePressureDataProfileRequest represents the request to delete a pressure data profile by ID.
type DeletePressureDataProfileRequest struct {
	ID string
}
