package cart

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/cart"
)

type CartRepository interface {
	GetAll(ctx context.Context, id uint) (cartOut []cart.Cart, err error)
	Create(ctx context.Context, cartIn cart.Cart) (err error)
	Update(ctx context.Context, cartIn cart.Cart) (err error)
	Delete(ctx context.Context, id uint) (err error)
}
