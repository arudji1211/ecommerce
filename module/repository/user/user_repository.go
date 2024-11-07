package user

import (
	"context"

	MDuser "github.com/arudji1211/ecommerce/module/model/user"
)

type UserRepository interface {
	GetAll(ctx context.Context) (UsersOut []MDuser.User, err error)
	GetUserByUsername(ctx context.Context, Username string) (UserOut MDuser.User, err error)
	Create(ctx context.Context, UserIn MDuser.User) (MDuser.User, error)
	Update(ctx context.Context, UserIn MDuser.User) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
