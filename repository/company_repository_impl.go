package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/domain"
)

type CompanyRepositoryImpl struct {
}

func NewCompanyRepository() CompanyRepository {
	return &CompanyRepositoryImpl{}
}

func (c CompanyRepositoryImpl) GetAll(ctx context.Context, db *sql.DB) []domain.Company {
	query := "SELECT id,name,created_at FROM companies"
	rows, err := db.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var companies []domain.Company
	for rows.Next() {
		company := domain.Company{}
		err := rows.Scan(&company.Id, &company.Name, &company.CreatedAt)
		helper.PanicIfError(err)
		companies = append(companies, company)
	}
	return companies
}

func (c CompanyRepositoryImpl) Find(ctx context.Context, db *sql.DB, id int) (domain.Company, error) {
	query := "SELECT id, name, created_at FROM companies WHERE id = ?"
	rows, err := db.QueryContext(ctx, query, id)
	helper.PanicIfError(err)
	defer rows.Close()

	company := domain.Company{}
	if rows.Next() {
		err := rows.Scan(&company.Id, &company.Name, &company.CreatedAt)
		helper.PanicIfError(err)
		return company, nil
	} else {
		return company, errors.New("company is not found")
	}
}

func (c CompanyRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, company domain.Company) domain.Company {
	currentDate := helper.GetCurrentTimestamp()
	query := "INSERT INTO companies(name, created_at) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, query, company.Name, currentDate)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	company.Id = int(id)
	company.CreatedAt = currentDate
	return company
}

func (c CompanyRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, company domain.Company, id int) domain.Company {
	query := "UPDATE companies SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, company.Name, id)
	helper.PanicIfError(err)

	company.Id = id
	return company
}

func (c CompanyRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, company domain.Company) {
	query := "DELETE FROM companies WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, company.Id)
	helper.PanicIfError(err)
}
