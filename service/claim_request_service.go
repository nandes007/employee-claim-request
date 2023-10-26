package service

import (
	"context"
	"github.com/nandes007/employee-claim-request/model/web/claim_request"
)

type ClaimRequestService interface {
	Create(ctx context.Context, r claim_request.Request, token string) claim_request.Response
	Update(ctx context.Context, r claim_request.Request, id int, token string) claim_request.Response
	UpdateStatus(ctx context.Context, r claim_request.ApproveRequest, id int, token string) claim_request.Response
	Delete(ctx context.Context, id int, token string)
	Find(ctx context.Context, id int, token string) claim_request.Response
	GetAll(ctx context.Context, token string) []claim_request.Response
}
