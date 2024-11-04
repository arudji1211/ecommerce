package user

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/user"
)

type UserRepository interface {
	GetAll(ctx context.Context) (UsersOut []user.User, err error)
	GetUserByUsername(ctx context.Context, Username string) (UserOut user.User, err error)
	Create(ctx context.Context, UserIn user.User) (user.User, error)
	Update(ctx context.Context, UserIn user.User) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
