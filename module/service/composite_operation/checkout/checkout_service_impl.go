package checkout

import (
	"net/http"

	MdCart "github.com/arudji1211/ecommerce/module/model/cart"
	"github.com/arudji1211/ecommerce/module/model/response"
	SPD "github.com/arudji1211/ecommerce/module/service/productdetail"
	ST "github.com/arudji1211/ecommerce/module/service/transaction"
	STD "github.com/arudji1211/ecommerce/module/service/transactiondetail"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CheckoutServiceImpl struct {
	db                 *gorm.DB
	Sproductdetail     SPD.ProductDetailService
	Stransaction       ST.TransactionService
	Stransactiondetail STD.TransactiondetailService
}

func (c *CheckoutServiceImpl) CheckoutFromCart(ctx *gin.Context, cartIn []MdCart.Cart) {
	//panic("not implemented") // TODO: Implement
	Tx := c.db.Begin()
	//check Product availability
	for i := 0; i < len(cartIn); i++ {
		//get product stock
		//non transaksional
		stock, err := c.Sproductdetail.GetByID(ctx, nil, cartIn[i].Product.ID)
		if err != nil {
			Tx.Rollback()
			//
			ctx.AbortWithStatusJSON(
				http.StatusNotFound,
				response.ResponseError{
					Code:    http.StatusNotFound,
					Message: "produk tidak ditemukan",
				},
			)
			return
		}

		if cartIn[i].Quantity >= stock.Stock {
			///product tidak mencukupi
			ctx.AbortWithStatusJSON(
				http.StatusNotAcceptable,
				response.ResponseError{
					Code:    http.StatusNotAcceptable,
					Message: "stok produk tidak mencukupi",
				},
			)
			return
		}
	}

	//kurangi stock barang
	for i := 0; i < len(cartIn); i++ {
		c.Sproductdetail.UpdateStock(ctx, Tx, cartIn[i].Product.ID, "-", cartIn[i].Quantity)
	}

}

func (c *CheckoutServiceImpl) QuickCheckout(ctx *gin.Context, cartIn MdCart.Cart) {
	panic("not implemented") // TODO: Implement
}

func NewCheckoutService(db *gorm.DB) CheckoutService {
	return &CheckoutServiceImpl{
		db: db,
	}
}
