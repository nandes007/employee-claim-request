package auth

type RegisterRequest struct {
	Name      string `validate:"required,max=100" json:"name"`
	CompanyId int    `json:"company_id"`
	Email     string `validate:"required,email,max=100" json:"email"`
	Password  string `validate:"required,min=8" json:"password"`
}
