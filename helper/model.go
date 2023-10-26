package helper

import (
	"github.com/nandes007/employee-claim-request/model/domain"
	"github.com/nandes007/employee-claim-request/model/web/claim_request"
	"github.com/nandes007/employee-claim-request/model/web/company"
	"github.com/nandes007/employee-claim-request/model/web/user"
)

func ToCompanyResponse(c domain.Company) company.Response {
	return company.Response{
		Id:        c.Id,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
	}
}

func ToCompanyResponses(c []domain.Company) []company.Response {
	var companies []company.Response
	for _, company := range c {
		companies = append(companies, ToCompanyResponse(company))
	}
	return companies
}

func ToClaimRequestResponse(c domain.ClaimRequest) claim_request.Response {
	return claim_request.Response{
		Id:              c.Id,
		UserId:          c.UserId,
		ClaimCategory:   c.ClaimCategory,
		ClaimDate:       c.ClaimDate,
		Currency:        c.Currency,
		ClaimAmount:     c.ClaimAmount,
		Description:     c.Description,
		SupportDocument: c.SupportDocument,
		Status:          c.Status,
		CreatedAt:       c.CreatedAt,
	}
}

func ToClaimRequestResponses(c []domain.ClaimRequest) []claim_request.Response {
	var claimRequests []claim_request.Response
	for _, claimRequest := range c {
		claimRequests = append(claimRequests, ToClaimRequestResponse(claimRequest))
	}
	return claimRequests
}

func ToUserResponse(u domain.User) user.Response {
	return user.Response{
		Id:        u.Id,
		Name:      u.Name,
		Email:     u.Email,
		CompanyId: u.CompanyId,
		CreatedAt: u.CreatedAt,
	}
}
