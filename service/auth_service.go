package service

import (
	"context"
	"github.com/nandes007/employee-claim-request/model/domain"
	"github.com/nandes007/employee-claim-request/model/web/auth"
)

type AuthService interface {
	Register(ctx context.Context, r auth.RegisterRequest, isAdmin bool) auth.RegisterResponse
	Login(ctx context.Context, user domain.User) string
}
