package user

import (
	UIModel "github.com/arudji1211/ecommerce/module/model/userinformation"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserInformation UIModel.UserInformation
	Username        string
	Password        string
	Role            string
}
