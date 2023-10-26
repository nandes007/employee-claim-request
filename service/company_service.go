package service

import (
	"context"
	"github.com/nandes007/employee-claim-request/model/web/company"
)

type CompanyService interface {
	Create(ctx context.Context, r company.Request) company.Response
	Update(ctx context.Context, r company.Request, companyId int) company.Response
	Delete(ctx context.Context, companyId int)
	Find(ctx context.Context, companyId int) company.Response
	GetAll(ctx context.Context) []company.Response
}
