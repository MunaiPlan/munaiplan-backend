package requests

// CreateStringRequestBody represents the request body for creating a String along with its Sections.
type CreateStringRequestBody struct {
	Name     string                     `json:"name"`
	Depth    float64                    `json:"depth"`
	Sections []CreateSectionRequestBody `json:"sections"`
}

// CreateSectionRequestBody represents the request body for creating a Section associated with a String.
type CreateSectionRequestBody struct {
	Description         *string  `json:"description,omitempty"`
	Manufacturer        *string  `json:"manufacturer,omitempty"`
	Type                string   `json:"type"`
	BodyMD              float64  `json:"body_md"`
	BodyLength          float64  `json:"body_length"`
	BodyOD              float64  `json:"body_od"`
	BodyID              float64  `json:"body_id"`
	AvgJointLength      *float64 `json:"avg_joint_length,omitempty"`
	StabilizerLength    *float64 `json:"stabilizer_length,omitempty"`
	StabilizerOD        *float64 `json:"stabilizer_od,omitempty"`
	StabilizerID        *float64 `json:"stabilizer_id,omitempty"`
	Weight              *float64 `json:"weight,omitempty"`
	Material            *string  `json:"material,omitempty"`
	Grade               *string  `json:"grade,omitempty"`
	Class               *int     `json:"class,omitempty"`
	FrictionCoefficient *float64 `json:"friction_coefficient,omitempty"`
	MinYieldStrength    *float64 `json:"min_yield_strength,omitempty"`
}

// CreateStringRequest represents the request for creating a String with Sections.
type CreateStringRequest struct {
	Body   CreateStringRequestBody
	CaseID string
}

// UpdateStringRequestBody represents the request body for updating a String and its associated Sections.
type UpdateStringRequestBody struct {
	Name     string                     `json:"name"`
	Depth    float64                    `json:"depth"`
	Sections []UpdateSectionRequestBody `json:"sections"`
}

// UpdateSectionRequestBody represents the request body for updating a Section associated with a String.
type UpdateSectionRequestBody struct {
	ID                  string   `json:"id"`
	Description         *string  `json:"description,omitempty"`
	Manufacturer        *string  `json:"manufacturer,omitempty"`
	Type                string   `json:"type"`
	BodyMD              float64  `json:"body_md"`
	BodyLength          float64  `json:"body_length"`
	BodyOD              float64  `json:"body_od"`
	BodyID              float64  `json:"body_id"`
	AvgJointLength      *float64 `json:"avg_joint_length,omitempty"`
	StabilizerLength    *float64 `json:"stabilizer_length,omitempty"`
	StabilizerOD        *float64 `json:"stabilizer_od,omitempty"`
	StabilizerID        *float64 `json:"stabilizer_id,omitempty"`
	Weight              *float64 `json:"weight,omitempty"`
	Material            *string  `json:"material,omitempty"`
	Grade               *string  `json:"grade,omitempty"`
	Class               *int     `json:"class,omitempty"`
	FrictionCoefficient *float64 `json:"friction_coefficient,omitempty"`
	MinYieldStrength    *float64 `json:"min_yield_strength,omitempty"`
}

// UpdateStringRequest represents the request for updating a String with its Sections.
type UpdateStringRequest struct {
	ID   string
	Body UpdateStringRequestBody
}

// GetStringsRequest represents the request for retrieving Strings by Case ID.
type GetStringsRequest struct {
	CaseID string
}

// GetStringByIDRequest represents the request for retrieving a String by its ID.
type GetStringByIDRequest struct {
	ID string
}

// DeleteStringRequest represents the request for deleting a String by its ID.
type DeleteStringRequest struct {
	ID string
}
