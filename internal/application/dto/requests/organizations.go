package requests

// CreateOrganizationRequest represents a request to create an organization.
type CreateOrganizationRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

// UpdateOrganizationRequest represents a request to update an organization.
type UpdateOrganizationRequest struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

// DeleteOrganizationRequest represents a request to delete an organization.
type DeleteOrganizationRequest struct {
	ID string `json:"id"`
}
