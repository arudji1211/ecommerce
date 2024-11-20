package userinformation

import (
	"context"

	MDuserinformation "github.com/arudji1211/ecommerce/module/model/userinformation"
	"gorm.io/gorm"
)

type UserInformationService interface {
	GetAll(ctx context.Context, db *gorm.DB) (UsersOut []MDuserinformation.UserInformation, err error)
	GetUserInformationByUserID(ctx context.Context, db *gorm.DB, UserID string) (UserOut MDuserinformation.UserInformation, err error)
	Create(ctx context.Context, db *gorm.DB, UserIn MDuserinformation.UserInformation) (MDuserinformation.UserInformation, error)
	Update(ctx context.Context, db *gorm.DB, UserIn MDuserinformation.UserInformation) (err error)
	Delete(ctx context.Context, db *gorm.DB, ID uint) (err error)
}
