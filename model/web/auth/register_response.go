package auth

type RegisterResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
