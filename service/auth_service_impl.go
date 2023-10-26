package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/nandes007/employee-claim-request/exception"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/domain"
	"github.com/nandes007/employee-claim-request/model/web/auth"
	"github.com/nandes007/employee-claim-request/repository"
)

type AuthServiceImpl struct {
	AuthRepository    repository.AuthRepository
	CompanyRepository repository.CompanyRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, companyRepository repository.CompanyRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository:    authRepository,
		CompanyRepository: companyRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request auth.RegisterRequest, isAdmin bool) auth.RegisterResponse {
	err := service.Validate.Struct(request)
	response := auth.RegisterResponse{}
	helper.PanicIfError(err)

	if request.CompanyId != 0 {
		_, err = service.CompanyRepository.Find(ctx, service.DB, request.CompanyId)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		CompanyId: request.CompanyId,
		IsAdmin:   isAdmin,
	}

	user, err = service.AuthRepository.Register(ctx, tx, user)
	helper.PanicIfError(err)

	response.Name = user.Name
	response.Email = user.Email

	return response
}

func (service *AuthServiceImpl) Login(ctx context.Context, user domain.User) string {
	generateToken, err := service.AuthRepository.Login(ctx, service.DB, user)

	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	return generateToken
}
