package requests

// CreateCaseRequestBody represents the request body for creating a case
type CreateCaseRequestBody struct {
	CaseName        string  `json:"case_name"`
	CaseDescription string  `json:"case_description"`
	DrillDepth      float64 `json:"drill_depth"`
	PipeSize        float64 `json:"pipe_size"`
}

// CreateCaseRequest represents the request for creating a case
type CreateCaseRequest struct {
	Body        CreateCaseRequestBody
	TrajectoryID string
}

// UpdateCaseRequestBody represents the request body for updating a case
type UpdateCaseRequestBody struct {
	CaseName        string  `json:"case_name"`
	CaseDescription string  `json:"case_description"`
	DrillDepth      float64 `json:"drill_depth"`
	PipeSize        float64 `json:"pipe_size"`
}

// UpdateCaseRequest represents the request for updating a case
type UpdateCaseRequest struct {
	ID   string
	Body UpdateCaseRequestBody
}

// GetCasesRequest represents the request for getting cases
type GetCasesRequest struct {
	TrajectoryID string
}

// GetCaseByIDRequest represents the request for getting a case by ID
type GetCaseByIDRequest struct {
	ID string
}

// DeleteCaseRequest represents the request for deleting a case
type DeleteCaseRequest struct {
	ID string
}
