package requests

type CreateCompanyRequestBody struct {
	Name           string `json:"name"`
	Division       string `json:"division"`
	Group          string `json:"group"`
	Representative string `json:"representative"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
}

type CreateCompanyRequest struct {
	OrganizationID string
	Body           CreateCompanyRequestBody
}

type UpdateCompanyRequest struct {
	OrganizationID string
	Body           UpdateCompanyRequestBody
}

type UpdateCompanyRequestBody struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Division       string `json:"division"`
	Group          string `json:"group"`
	Representative string `json:"representative"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
}

type GetCompaniesRequest struct {
	OrganizationID string
}

type GetCompanyByNameRequest struct {
	OrganizationID string
	Name           string
}

type GetCompanyByID struct {
	OrganizationID string
	ID             string
}

type DeleteCompanyRequest struct {
	OrganizationID string
	ID             string
}
