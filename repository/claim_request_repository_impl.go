package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/domain"
)

type ClaimRequestRepositoryImpl struct {
}

func NewClaimRequestRepository() ClaimRequestRepository {
	return &ClaimRequestRepositoryImpl{}
}

func (c ClaimRequestRepositoryImpl) GetAll(ctx context.Context, db *sql.DB, user domain.User) []domain.ClaimRequest {
	var (
		query    string
		identify int
	)
	if user.IsAdmin {
		query = `SELECT cl.id,cl.user_id,cl.claim_category,cl.claim_date,cl.currency,cl.claim_amount,cl.description,
			cl.support_document,cl.status,cl.created_at FROM claim_requests AS cl
		LEFT JOIN users AS u ON cl.user_id = u.id WHERE u.company_id = ?`
		identify = user.CompanyId
	} else {
		query = "SELECT id,user_id,claim_category,claim_date,currency,claim_amount,description,support_document,status,created_at FROM claim_requests WHERE user_id = ?"
		identify = user.Id
	}

	rows, err := db.QueryContext(ctx, query, identify)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var claimRequests []domain.ClaimRequest
	for rows.Next() {
		claimRequest := domain.ClaimRequest{}
		err := rows.Scan(&claimRequest.Id, &claimRequest.UserId, &claimRequest.ClaimCategory, &claimRequest.ClaimDate, &claimRequest.Currency, &claimRequest.ClaimAmount, &claimRequest.Description, &claimRequest.SupportDocument, &claimRequest.Status, &claimRequest.CreatedAt)
		helper.PanicIfError(err)
		claimRequests = append(claimRequests, claimRequest)
	}

	return claimRequests
}

func (c ClaimRequestRepositoryImpl) Find(ctx context.Context, db *sql.DB, user domain.User, id int) (domain.ClaimRequest, error) {
	var (
		query    string
		identify int
	)
	if user.IsAdmin {
		query = `SELECT cl.id,cl.user_id,cl.claim_category,cl.claim_date,cl.currency,cl.claim_amount,cl.description,
			cl.support_document,cl.status,cl.created_at FROM claim_requests AS cl
		LEFT JOIN users AS u ON cl.user_id = u.id WHERE cl.id = ? AND u.company_id = ?`
		identify = user.CompanyId
	} else {
		query = "SELECT id,user_id,claim_category,claim_date,currency,claim_amount,description,support_document,status,created_at FROM claim_requests WHERE id = ? AND user_id = ?"
		identify = user.Id
	}
	rows, err := db.QueryContext(ctx, query, id, identify)
	helper.PanicIfError(err)
	defer rows.Close()

	claimRequest := domain.ClaimRequest{}
	if rows.Next() {
		err := rows.Scan(&claimRequest.Id, &claimRequest.UserId, &claimRequest.ClaimCategory, &claimRequest.ClaimDate, &claimRequest.Currency, &claimRequest.ClaimAmount, &claimRequest.Description, &claimRequest.SupportDocument, &claimRequest.Status, &claimRequest.CreatedAt)
		helper.PanicIfError(err)
		return claimRequest, nil
	} else {
		return claimRequest, errors.New("claim request is not found")
	}
}

func (c ClaimRequestRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, claimRequest domain.ClaimRequest) domain.ClaimRequest {
	currentDate := helper.GetCurrentTimestamp()
	query := "INSERT INTO claim_requests(user_id, claim_category, claim_date, currency, claim_amount, description, support_document, status, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, claimRequest.UserId, claimRequest.ClaimCategory, claimRequest.ClaimDate, claimRequest.Currency, claimRequest.ClaimAmount, claimRequest.Description, claimRequest.SupportDocument, claimRequest.Status, currentDate)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	claimRequest.Id = int(id)
	claimRequest.CreatedAt = currentDate
	return claimRequest
}

func (c ClaimRequestRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, claimRequest domain.ClaimRequest, id int) domain.ClaimRequest {
	query := "UPDATE claim_requests SET user_id = ?, claim_category = ?, claim_date = ?, currency = ?, claim_amount = ?, description = ?, support_document = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, claimRequest.UserId, claimRequest.ClaimCategory, claimRequest.ClaimDate, claimRequest.Currency, claimRequest.ClaimAmount, claimRequest.Description, claimRequest.SupportDocument, id)
	helper.PanicIfError(err)

	claimRequest.Id = id
	return claimRequest
}

func (c ClaimRequestRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, claimRequest domain.ClaimRequest, user domain.User) domain.ClaimRequest {
	query := "UPDATE claim_requests SET status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, claimRequest.Status, claimRequest.Id)
	helper.PanicIfError(err)
	return claimRequest
}

func (c ClaimRequestRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, claimRequest domain.ClaimRequest) {
	query := "DELETE FROM claim_requests WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, claimRequest.Id)
	helper.PanicIfError(err)
}
