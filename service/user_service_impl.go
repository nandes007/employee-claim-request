package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/nandes007/employee-claim-request/exception"
	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/web/user"
	"github.com/nandes007/employee-claim-request/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(repository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: repository,
		DB:             DB,
		Validate:       validate,
	}
}

func (u UserServiceImpl) Update(ctx context.Context, r user.Request, token string) user.Response {
	err := u.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, err := u.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := u.UserRepository.Find(ctx, u.DB, token)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Name = r.Name
	user.CompanyId = r.CompanyId

	user = u.UserRepository.Update(ctx, tx, user)
	return helper.ToUserResponse(user)
}
