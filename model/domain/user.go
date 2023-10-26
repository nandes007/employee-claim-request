package domain

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CompanyId int    `json:"company_id"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt string `json:"created_at"`
}
