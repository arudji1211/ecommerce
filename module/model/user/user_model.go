package user

import (
	"github.com/arudji1211/ecommerce/module/model/cart"
	UIModel "github.com/arudji1211/ecommerce/module/model/userinformation"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserDetail UIModel.UserInformation
	Username   string
	Password   string
	Role       string
	Cart       []cart.Cart
}
