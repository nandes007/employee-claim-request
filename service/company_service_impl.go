package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/nandes007/employee-claim-request/exception"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/domain"
	"github.com/nandes007/employee-claim-request/model/web/company"
	"github.com/nandes007/employee-claim-request/repository"
)

type CompanyServiceImpl struct {
	CompanyRepository repository.CompanyRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCompanyService(repository repository.CompanyRepository, DB *sql.DB, validate *validator.Validate) CompanyService {
	return &CompanyServiceImpl{
		CompanyRepository: repository,
		DB:                DB,
		Validate:          validate,
	}
}

func (c CompanyServiceImpl) Create(ctx context.Context, r company.Request) company.Response {
	err := c.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	company := domain.Company{
		Name: r.Name,
	}

	company = c.CompanyRepository.Save(ctx, tx, company)
	return helper.ToCompanyResponse(company)
}

func (c CompanyServiceImpl) Update(ctx context.Context, r company.Request, companyId int) company.Response {
	err := c.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	company, err := c.CompanyRepository.Find(ctx, c.DB, companyId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	company.Name = r.Name

	company = c.CompanyRepository.Update(ctx, tx, company, companyId)

	return helper.ToCompanyResponse(company)
}

func (c CompanyServiceImpl) Delete(ctx context.Context, companyId int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	company, err := c.CompanyRepository.Find(ctx, c.DB, companyId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	c.CompanyRepository.Delete(ctx, tx, company)
}

func (c CompanyServiceImpl) Find(ctx context.Context, companyId int) company.Response {
	company, err := c.CompanyRepository.Find(ctx, c.DB, companyId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCompanyResponse(company)
}

func (c CompanyServiceImpl) GetAll(ctx context.Context) []company.Response {
	companies := c.CompanyRepository.GetAll(ctx, c.DB)
	return helper.ToCompanyResponses(companies)
}
