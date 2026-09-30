package requests

type AdminUserInput struct {
	Name     string `json:"name" binding:"required,max=64"`
	Surname  string `json:"surname" binding:"required,max=64"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Phone    string `json:"phone" binding:"omitempty,max=20"`
	Password string `json:"password" binding:"required,max=64"`
	Role     string `json:"role" binding:"omitempty,oneof=user admin"`
}

type AdminOrganizationInput struct {
	Name    string `json:"name" binding:"required,max=255"`
	Email   string `json:"email" binding:"required,email,max=255"`
	Phone   string `json:"phone" binding:"omitempty,max=20"`
	Address string `json:"address" binding:"omitempty,max=500"`
}

// AdminCreateOrganizationRequest creates a tenant together with its first account.
type AdminCreateOrganizationRequest struct {
	Organization AdminOrganizationInput `json:"organization" binding:"required"`
	User         AdminUserInput         `json:"user" binding:"required"`
}
