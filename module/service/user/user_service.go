package user

import (
	"context"

	USmodel "github.com/arudji1211/ecommerce/module/model/user"
	"gorm.io/gorm"
)

type UserService interface {
	GetAll(ctx context.Context, db *gorm.DB) (userOut []USmodel.User, err error)
	GetByUsername(ctx context.Context, db *gorm.DB, username string) (userOut USmodel.User, err error)
	Create(ctx context.Context, db *gorm.DB, userIn USmodel.User) (userOut USmodel.User, err error)
	Update(ctx context.Context, db *gorm.DB, userIn USmodel.User) (err error)
	Delete(ctx context.Context, db *gorm.DB, ID uint) (err error)
}
