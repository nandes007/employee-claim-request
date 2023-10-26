package repository

import (
	"context"
	"database/sql"
	"github.com/nandes007/employee-claim-request/model/domain"
)

type ClaimRequestRepository interface {
	GetAll(ctx context.Context, db *sql.DB, user domain.User) []domain.ClaimRequest
	Find(ctx context.Context, db *sql.DB, user domain.User, id int) (domain.ClaimRequest, error)
	Save(ctx context.Context, tx *sql.Tx, claimRequest domain.ClaimRequest) domain.ClaimRequest
	Update(ctx context.Context, tx *sql.Tx, claimRequest domain.ClaimRequest, id int) domain.ClaimRequest
	UpdateStatus(ctx context.Context, tx *sql.Tx, claimRequest domain.ClaimRequest, user domain.User) domain.ClaimRequest
	Delete(ctx context.Context, tx *sql.Tx, claimRequest domain.ClaimRequest)
}
