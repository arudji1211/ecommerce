package cart

import (
	"context"

	Mcart "github.com/arudji1211/ecommerce/module/model/cart"
)

type CartService interface {
	GetAll(ctx context.Context, id uint) (cartOut []Mcart.Cart, err error)
	Create(ctx context.Context, cartIn Mcart.Cart) (Mcart.Cart, error)
	Update(ctx context.Context, cartIn Mcart.Cart) (err error)
	Delete(ctx context.Context, id uint) (err error)
}
