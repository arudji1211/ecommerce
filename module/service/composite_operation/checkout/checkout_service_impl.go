package checkout

import (
	"net/http"

	MdCart "github.com/arudji1211/ecommerce/module/model/cart"
	"github.com/arudji1211/ecommerce/module/model/response"
	"github.com/arudji1211/ecommerce/module/model/transaction"
	MdTransaction "github.com/arudji1211/ecommerce/module/model/transaction"
	MdTransDetail "github.com/arudji1211/ecommerce/module/model/transactionDetail"
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

func (c *CheckoutServiceImpl) CheckoutFromCart(ctx *gin.Context, ID uint, cartIn []MdCart.Cart) (TransactionOut transaction.Transaction, err error) {
	//panic("not implemented") // TODO: Implement
	Tx := c.db.Begin()
	//check Product availability
	for i := 0; i < len(cartIn); i++ {
		//get product stock
		//non transaksional
		stock, errs := c.Sproductdetail.GetByID(ctx, nil, cartIn[i].Product.ID)
		if errs != nil {
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

	//buat transaksi
	TransactionOut, err = c.QuickCheckout(ctx, ID, Tx, cartIn)
	if err != nil {
		Tx.Rollback()
		//
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Code:    http.StatusInternalServerError,
				Message: "Gagal Checkout",
			},
		)
		return
	}

	//kurangi stock barang
	for i := 0; i < len(cartIn); i++ {
		err := c.Sproductdetail.UpdateStock(ctx, Tx, cartIn[i].Product.ID, "-", cartIn[i].Quantity)
		if err != nil {
			Tx.Rollback()
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				response.ResponseError{
					Code:    http.StatusInternalServerError,
					Message: "Kesalahan saat mengakses data " + cartIn[i].Product.Title,
				},
			)
		}
	}

	return

}

func (c *CheckoutServiceImpl) QuickCheckout(ctx *gin.Context, ID uint, tx *gorm.DB, cartIn []MdCart.Cart) (TransactionOut transaction.Transaction, err error) {
	// panic("not implemented") // TODO: Implement

	// buat transaksi
	// total harga
	var amount int = 0

	for i := 0; i < len(cartIn); i++ {
		amount += cartIn[i].Quantity * cartIn[i].ProductDetail.Harga
	}
	//cost tambahan
	//ex biaya penanganan 12% dari amount
	amount += 15000

	//payment method
	payment_method := "CASH ON DELIVERY"

	//status
	Status_pesanan := "Belum Bayar"

	//buat wrapper transaksi
	TransactionOut = MdTransaction.Transaction{
		UserID:        ID,
		PaymentMethod: payment_method,
		Status:        Status_pesanan,
		Amount:        amount,
	}
	TransactionOut, err = c.Stransaction.Create(ctx, tx, ID, TransactionOut)
	if err != nil {
		tx.Rollback()
		//
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			response.ResponseError{
				Code:    http.StatusInternalServerError,
				Message: "Gagal Membuat Transaksi",
			},
		)
		return
	}

	///buat transaksi detail
	for i := 0; i < len(cartIn); i++ {
		transaksiDetail := MdTransDetail.TransactionDetail{
			TransactionID:   TransactionOut.ID,
			ProductDetailID: cartIn[i].ProductDetailID,
			Quantity:        cartIn[i].Quantity,
		}
		//panggil service
		transaksiDetail, errs := c.Stransactiondetail.Create(ctx, tx, transaksiDetail)
		if errs != nil {
			tx.Rollback()
			//
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				response.ResponseError{
					Code:    http.StatusInternalServerError,
					Message: "Gagal Membuat Transaksi Detail",
				},
			)
			return
		}
	}

	//berhasil checkout
	return
}

func NewCheckoutService(db *gorm.DB) CheckoutService {
	return &CheckoutServiceImpl{
		db: db,
	}
}
