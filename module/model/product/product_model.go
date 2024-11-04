package product

import (
	productdetails "github.com/arudji1211/ecommerce/module/model/productDetail"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	Description string
	Detail      productdetails.ProductDetail
}
