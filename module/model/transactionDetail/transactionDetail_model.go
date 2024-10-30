package transactiondetail

import (
	productdetail "github.com/arudji1211/ecommerce/module/model/productDetail"
	"gorm.io/gorm"
)

type TransactionDetail struct {
	gorm.Model
	TransactionID   uint
	ProductDetailID uint
	Quantity        int
	ProductDetail   productdetail.ProductDetail
}
