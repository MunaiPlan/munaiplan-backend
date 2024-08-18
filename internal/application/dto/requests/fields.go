package requests

// CreateFieldRequestBody represents the request body for creating a field
type CreateFieldRequestBody struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ReductionLevel  string `json:"reduction_level"`
	ActiveFieldUnit string `json:"active_field_unit"`
}

// CreateFieldRequest represents the request for creating a field
type CreateFieldRequest struct {
	Body      CreateFieldRequestBody
	CompanyID string
}

// UpdateFieldRequestBody represents the request body for updating a field
type UpdateFieldRequestBody struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ReductionLevel  string `json:"reduction_level"`
	ActiveFieldUnit string `json:"active_field_unit"`
}

// UpdateFieldRequest represents the request for updating a field
type UpdateFieldRequest struct {
	ID   string
	Body UpdateFieldRequestBody
}

// GetFieldsRequest represents the request for getting fields
type GetFieldsRequest struct {
	CompanyID string
}

// GetFieldByIDRequest represents the request for getting a field by ID
type GetFieldByIDRequest struct {
	ID        string
}

// DeleteFieldRequest represents the request for deleting a field
type DeleteFieldRequest struct {
	ID        string
}
