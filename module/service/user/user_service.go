package user

import (
	"context"

	USmodel "github.com/arudji1211/ecommerce/module/model/user"
)

type UserService interface {
	GetAll(ctx context.Context) (userOut []USmodel.User, err error)
	GetByUsername(ctx context.Context, username string) (userOut USmodel.User, err error)
	Create(ctx context.Context, userIn USmodel.User) (userOut USmodel.User, err error)
	Update(ctx context.Context, userIn USmodel.User) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
