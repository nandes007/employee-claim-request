package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/nandes007/employee-claim-request/helper"
	"github.com/nandes007/employee-claim-request/model/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	currentDate := helper.GetCurrentTimestamp()
	passwordHashed := helper.PasswordHash(user.Password)
	var nullableCompany *int
	if user.CompanyId == 0 {
		nullableCompany = nil
	} else {
		nullableCompany = new(int)
		*nullableCompany = user.CompanyId
	}
	stmt := "INSERT INTO users(name, email, password, company_id, is_admin, created_at) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, stmt, user.Name, user.Email, passwordHashed, nullableCompany, user.IsAdmin, currentDate)

	if err != nil {
		fmt.Println("Here error 1")
		log.Fatal(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println("Here error 2")
		log.Fatal(err)
	}

	user.Id = int(id)
	return user, nil
}

func (repository AuthRepositoryImpl) Login(ctx context.Context, db *sql.DB, user domain.User) (string, error) {
	var (
		id       int
		email    string
		password string
	)
	stmt, err := db.PrepareContext(ctx, "SELECT id, email, password FROM users WHERE email = ? LIMIT 1")
	helper.PanicIfError(err)
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, user.Email).Scan(&id, &email, &password)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user is not found")
		} else {
			helper.PanicIfError(err)
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
	if err != nil {
		return "", errors.New("credential mismatch")
	}

	tokenString, err := helper.CreateToken(id)
	helper.PanicIfError(err)

	return tokenString, nil
}
