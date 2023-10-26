package claim_request

type ApproveRequest struct {
	Status string `validate:"required", json:"status"`
}
