package cart

import (
	"github.com/arudji1211/ecommerce/module/model/product"
	"github.com/arudji1211/ecommerce/module/model/productdetail"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID          uint
	ProductID       uint
	Product         product.Product `json:"product"`
	ProductDetailID uint
	ProductDetail   productdetail.ProductDetail `json:"product_detail"`
	Quantity        int
}
