package claim_request

type Request struct {
	ClaimCategory   string  `validate:"required" json:"claim_category"`
	ClaimDate       string  `validate:"required" json:"claim_date"`
	Currency        string  `validate:"required" json:"currency"`
	ClaimAmount     float64 `validate:"required,numeric" json:"claim_amount"`
	Description     string  `validate:"required" json:"description"`
	SupportDocument string  `validate:"required" json:"support_document"`
}
