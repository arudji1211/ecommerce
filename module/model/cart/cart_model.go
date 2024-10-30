package cart

import (
	"github.com/arudji1211/ecommerce/module/model/productdetail"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID          uint
	ProductDetailID uint
	Quantity        int
	Product         productdetail.ProductDetail `json:"Product"`
}
