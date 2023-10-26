package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Find(ctx context.Context, db *sql.DB, token string) (domain.User, error) {
	userId, err := helper.ParseUserToken(token)

	if err != nil {
		return domain.User{}, errors.New("credential mismatch")
	}

	var user domain.User
	stmt, err := db.PrepareContext(ctx, "SELECT id, email, ifnull(company_id, 0), is_admin FROM users WHERE id = ? LIMIT 1")
	helper.PanicIfError(err)
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, userId).Scan(&user.Id, &user.Email, &user.CompanyId, &user.IsAdmin)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("user is not found")
		} else {
			log.Fatal(err)
			helper.PanicIfError(err)
		}
	}

	return user, nil
}

func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	query := "UPDATE users SET name = ?, company_id = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Name, user.CompanyId, user.Id)
	helper.PanicIfError(err)

	return user
}
