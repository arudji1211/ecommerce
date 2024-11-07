package userinformation

import (
	"context"

	MDuserinformation "github.com/arudji1211/ecommerce/module/model/userinformation"
)

type UserInformationRepository interface {
	GetAll(ctx context.Context) (UsersOut []MDuserinformation.UserInformation, err error)
	GetUserInformationByUserID(ctx context.Context, UserID string) (UserOut MDuserinformation.UserInformation, err error)
	Create(ctx context.Context, UserIn MDuserinformation.UserInformation) (MDuserinformation.UserInformation, error)
	Update(ctx context.Context, UserIn MDuserinformation.UserInformation) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
