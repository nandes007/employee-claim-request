package user

type Response struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CompanyId int    `json:"companyId"`
	CreatedAt string `json:"created_at"`
}
