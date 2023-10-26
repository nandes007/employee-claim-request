package repository

import (
	"context"
	"database/sql"
	"github.com/nandes007/employee-claim-request/model/domain"
)

type UserRepository interface {
	Find(ctx context.Context, db *sql.DB, token string) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
}
