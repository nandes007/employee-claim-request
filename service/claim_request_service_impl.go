package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nandes007/employee-claim-request/exception"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/domain"
	"github.com/nandes007/employee-claim-request/model/web/claim_request"
	"github.com/nandes007/employee-claim-request/repository"
)

type ClaimRequestServiceImpl struct {
	Repository     repository.ClaimRequestRepository
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewClaimRequestService(repository repository.ClaimRequestRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) ClaimRequestService {
	return &ClaimRequestServiceImpl{
		Repository:     repository,
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (c ClaimRequestServiceImpl) Create(ctx context.Context, r claim_request.Request, token string) claim_request.Response {
	err := c.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tokenFormatted := helper.FormatToken(token)
	user, err := c.UserRepository.Find(ctx, c.DB, tokenFormatted)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	if user.CompanyId == 0 {
		err = errors.New("sorry, you don't have company")
		panic(exception.NewBadRequestError(err.Error()))
	}

	if user.IsAdmin {
		err = errors.New("sorry, you don't have access to this feature")
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	_, err = time.Parse("2006-01-02", r.ClaimDate)

	if err != nil {
		err = errors.New("sorry, claim date using invalid format")
		panic(exception.NewBadRequestError(err.Error()))
	}

	claimRequest := domain.ClaimRequest{
		UserId:          user.Id,
		ClaimCategory:   r.ClaimCategory,
		ClaimDate:       r.ClaimDate,
		Currency:        r.Currency,
		ClaimAmount:     r.ClaimAmount,
		Description:     r.Description,
		SupportDocument: r.SupportDocument,
		Status:          "Pending",
	}

	claimRequest = c.Repository.Save(ctx, tx, claimRequest)
	return helper.ToClaimRequestResponse(claimRequest)
}

func (c ClaimRequestServiceImpl) Update(ctx context.Context, r claim_request.Request, id int, token string) claim_request.Response {
	err := c.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tokenFormatted := helper.FormatToken(token)
	user, err := c.UserRepository.Find(ctx, c.DB, tokenFormatted)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	if user.CompanyId == 0 {
		err = errors.New("sorry, you don't have company")
		panic(exception.NewBadRequestError(err.Error()))
	}

	if user.IsAdmin {
		err = errors.New("sorry, you don't have access to this feature")
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	claimRequest, err := c.Repository.Find(ctx, c.DB, user, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	_, err = time.Parse("2006-01-02", r.ClaimDate)

	if err != nil {
		err = errors.New("sorry, claim date using invalid format")
		panic(exception.NewBadRequestError(err.Error()))
	}

	claimRequest.UserId = user.Id
	claimRequest.ClaimCategory = r.ClaimCategory
	claimRequest.ClaimDate = r.ClaimDate
	claimRequest.Currency = r.Currency
	claimRequest.ClaimAmount = r.ClaimAmount
	claimRequest.Description = r.Description
	claimRequest.SupportDocument = r.SupportDocument
	claimRequest.Status = "Pending"

	claimRequest = c.Repository.Update(ctx, tx, claimRequest, id)
	return helper.ToClaimRequestResponse(claimRequest)
}

func (c ClaimRequestServiceImpl) UpdateStatus(ctx context.Context, r claim_request.ApproveRequest, id int, token string) claim_request.Response {
	err := c.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tokenFormatted := helper.FormatToken(token)
	user, err := c.UserRepository.Find(ctx, c.DB, tokenFormatted)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if !user.IsAdmin {
		err = errors.New("sorry, you don't have access to this feature")
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	claimRequest, err := c.Repository.Find(ctx, c.DB, user, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if r.Status != "Approved" || r.Status != "Rejected" {
		err = errors.New("you just can approve (Approved) or reject (Rejected) request")
		panic(exception.NewBadRequestError(err.Error()))
	}

	if claimRequest.Status == "Approved" || claimRequest.Status == "Rejected" {
		err = errors.New("this request has been " + claimRequest.Status)
		panic(exception.NewBadRequestError(err.Error()))
	}

	claimRequest.Status = r.Status
	claimRequest = c.Repository.UpdateStatus(ctx, tx, claimRequest, user)
	return helper.ToClaimRequestResponse(claimRequest)
}

func (c ClaimRequestServiceImpl) Delete(ctx context.Context, id int, token string) {
	tokenFormatted := helper.FormatToken(token)
	user, err := c.UserRepository.Find(ctx, c.DB, tokenFormatted)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	if user.CompanyId == 0 {
		err = errors.New("sorry, you don't have company")
		panic(exception.NewBadRequestError(err.Error()))
	}

	if user.IsAdmin {
		err = errors.New("sorry, you don't have access to this feature")
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	claimRequest, err := c.Repository.Find(ctx, c.DB, user, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	c.Repository.Delete(ctx, tx, claimRequest)
}

func (c ClaimRequestServiceImpl) Find(ctx context.Context, id int, token string) claim_request.Response {
	tokenFormatted := helper.FormatToken(token)
	user, err := c.UserRepository.Find(ctx, c.DB, tokenFormatted)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	claimRequest, err := c.Repository.Find(ctx, c.DB, user, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToClaimRequestResponse(claimRequest)
}

func (c ClaimRequestServiceImpl) GetAll(ctx context.Context, token string) []claim_request.Response {
	tokenFormatted := helper.FormatToken(token)
	user, err := c.UserRepository.Find(ctx, c.DB, tokenFormatted)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	claimRequests := c.Repository.GetAll(ctx, c.DB, user)
	return helper.ToClaimRequestResponses(claimRequests)
}
