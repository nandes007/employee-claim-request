package repository

import (
	"context"
	"database/sql"
	"github.com/nandes007/employee-claim-request/model/domain"
)

type CompanyRepository interface {
	GetAll(ctx context.Context, db *sql.DB) []domain.Company
	Find(ctx context.Context, db *sql.DB, id int) (domain.Company, error)
	Save(ctx context.Context, tx *sql.Tx, company domain.Company) domain.Company
	Update(ctx context.Context, tx *sql.Tx, company domain.Company, id int) domain.Company
	Delete(ctx context.Context, tx *sql.Tx, company domain.Company)
}
