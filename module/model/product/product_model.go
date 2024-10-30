package product

import (
	productdetail "github.com/arudji1211/ecommerce/module/model/productDetail"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	Description string
	Detail      productdetail.ProductDetail
}
