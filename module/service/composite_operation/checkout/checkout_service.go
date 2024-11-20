package checkout

import (
	"context"

	MdCart "github.com/arudji1211/ecommerce/module/model/cart"
)

type CheckoutService interface {
	CheckoutFromCart(ctx context.Context, cartIn []MdCart.Cart)
	QuickCheckout(ctx context.Context, cartIn MdCart.Cart)
}
