package product

import (
	productdetails "github.com/arudji1211/ecommerce/module/model/productdetail"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string
	Description string
	Detail      productdetails.ProductDetail
}

type ProductWithDetail struct {
	Product
	ProductDetail productdetails.ProductDetail
}
