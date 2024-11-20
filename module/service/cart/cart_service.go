package cart

import (
	"context"

	Mcart "github.com/arudji1211/ecommerce/module/model/cart"
	"gorm.io/gorm"
)

type CartService interface {
	GetAll(ctx context.Context, db *gorm.DB, id uint) (cartOut []Mcart.Cart, err error)
	Create(ctx context.Context, db *gorm.DB, cartIn Mcart.Cart) (Mcart.Cart, error)
	Update(ctx context.Context, db *gorm.DB, cartIn Mcart.Cart) (err error)
	Delete(ctx context.Context, db *gorm.DB, id uint) (err error)
}
