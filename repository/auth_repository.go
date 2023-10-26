package repository

import (
	"context"
	"database/sql"
	"github.com/nandes007/employee-claim-request/model/domain"
)

type AuthRepository interface {
	Register(ctx context.Context, db *sql.Tx, user domain.User) (domain.User, error)
	Login(ctx context.Context, db *sql.DB, user domain.User) (string, error)
}
