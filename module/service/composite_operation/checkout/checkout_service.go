package checkout

import (
	MdCart "github.com/arudji1211/ecommerce/module/model/cart"
	"github.com/arudji1211/ecommerce/module/model/transaction"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CheckoutService interface {
	CheckoutFromCart(ctx *gin.Context, ID uint, cartIn []MdCart.Cart) (TransactionOut transaction.Transaction, err error)
	QuickCheckout(ctx *gin.Context, ID uint, tx *gorm.DB, cartIn []MdCart.Cart) (TransactionOut transaction.Transaction, err error)
}
