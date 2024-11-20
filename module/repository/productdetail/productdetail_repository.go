package productdetail

import (
	"context"

	MDproductdetail "github.com/arudji1211/ecommerce/module/model/productdetail"
	"gorm.io/gorm"
)

type ProductDetailRepository interface {
	GetAllByProductID(ctx context.Context, db *gorm.DB, ID uint) (productdetailIn []MDproductdetail.ProductDetail, err error)
	GetByID(ctx context.Context, db *gorm.DB, ID uint) (productdetailIn MDproductdetail.ProductDetail, err error)
	Create(ctx context.Context, db *gorm.DB, productdetailIn MDproductdetail.ProductDetail) (MDproductdetail.ProductDetail, error)
	Update(ctx context.Context, db *gorm.DB, productdetailIn MDproductdetail.ProductDetail) (err error)
	Delete(ctx context.Context, db *gorm.DB, ID uint) (err error)
	IsTransaction(Db *gorm.DB)
}
