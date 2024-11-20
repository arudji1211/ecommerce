package user

import (
	"context"

	MDuser "github.com/arudji1211/ecommerce/module/model/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll(ctx context.Context, db *gorm.DB) (UsersOut []MDuser.User, err error)
	GetUserByUsername(ctx context.Context, db *gorm.DB, Username string) (UserOut MDuser.User, err error)
	Create(ctx context.Context, db *gorm.DB, UserIn MDuser.User) (MDuser.User, error)
	Update(ctx context.Context, db *gorm.DB, UserIn MDuser.User) (err error)
	Delete(ctx context.Context, db *gorm.DB, ID uint) (err error)
	IsTransaction(Db *gorm.DB)
}
