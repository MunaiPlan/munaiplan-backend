package requests

type UserSignInRequest struct {
	Body           UserSignInRequestBody
	OrganizationID string
}

type UserSignInRequestBody struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type UserSignUpRequest struct {
	Body           UserSignUpRequestBody
	OrganizationID string
}

type UserSignUpRequestBody struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
	Name     string `json:"name" binding:"required,min=3,max=32"`
	Surname  string `json:"surname" binding:"required,min=3,max=32"`
	Phone    string `json:"phone" binding:"omitempty,min=10,max=20"`
}
