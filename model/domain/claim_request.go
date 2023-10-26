package domain

type ClaimRequest struct {
	Id              int     `json:"id"`
	UserId          int     `json:"user_id"`
	ClaimCategory   string  `json:"claim_category"`
	ClaimDate       string  `json:"claim_date"`
	Currency        string  `json:"currency"`
	ClaimAmount     float64 `json:"claim_amount"`
	Description     string  `json:"description"`
	SupportDocument string  `json:"support_document"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"created_at"`
}
