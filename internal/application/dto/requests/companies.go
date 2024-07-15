package requests

type CreateCompanyRequest struct {
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	Division       string `json:"division"`
	Group          string `json:"group"`
	Representative string `json:"representative"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
}

type UpdateCompanyRequest struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Division       string `json:"division"`
	Group          string `json:"group"`
	Representative string `json:"representative"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
}

type DeleteCompanyRequest struct {
	ID string `json:"id"`
}
