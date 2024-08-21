package requests

import (
	"time"
)

// CreateDesignRequestBody represents the request body for creating a design
type CreateDesignRequestBody struct {
	PlanName   string    `json:"plan_name"`
	Stage      string    `json:"stage"`
	Version    string    `json:"version"`
	ActualDate time.Time `json:"actual_date"`
}

// CreateDesignRequest represents the request for creating a design
type CreateDesignRequest struct {
	Body       CreateDesignRequestBody
	WellboreID string
}

// UpdateDesignRequestBody represents the request body for updating a design
type UpdateDesignRequestBody struct {
	PlanName   string    `json:"plan_name"`
	Stage      string    `json:"stage"`
	Version    string    `json:"version"`
	ActualDate time.Time `json:"actual_date"`
}

// UpdateDesignRequest represents the request for updating a design
type UpdateDesignRequest struct {
	ID   string
	Body UpdateDesignRequestBody
}

// GetDesignsRequest represents the request for getting designs
type GetDesignsRequest struct {
	WellboreID string
}

// GetDesignByIDRequest represents the request for getting a design by ID
type GetDesignByIDRequest struct {
	ID string
}

// DeleteDesignRequest represents the request for deleting a design
type DeleteDesignRequest struct {
	ID string
}
