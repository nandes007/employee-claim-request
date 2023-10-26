package user

type Request struct {
	Name      string `validate:"required,max=100" json:"name"`
	CompanyId int    `validate:"required,number" json:"company_id"`
}
