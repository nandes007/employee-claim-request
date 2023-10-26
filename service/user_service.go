package service

import (
	"context"
	"github.com/nandes007/employee-claim-request/model/web/user"
)

type UserService interface {
	Update(ctx context.Context, r user.Request, token string) user.Response
}
