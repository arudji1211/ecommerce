package userinformation

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/userinformation"
)

type UserInformationRepository interface {
	GetAll(ctx context.Context) (UsersOut []userinformation.UserInformation, err error)
	GetUserInformationByUserID(ctx context.Context, UserID string) (UserOut userinformation.UserInformation, err error)
	Create(ctx context.Context, UserIn userinformation.UserInformation) (err error)
	Update(ctx context.Context, UserIn userinformation.UserInformation) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
