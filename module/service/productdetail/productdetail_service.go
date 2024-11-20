package productdetail

import (
	"context"

	PDmodel "github.com/arudji1211/ecommerce/module/model/productdetail"
	"gorm.io/gorm"
)

type ProductDetailService interface {
	GetAll(ctx context.Context, db *gorm.DB, ID uint) (productdetails []PDmodel.ProductDetail, err error)
	GetByID(ctx context.Context, db *gorm.DB, ID uint) (productdetail PDmodel.ProductDetail, err error)
	Create(ctx context.Context, db *gorm.DB, productdetailIn PDmodel.ProductDetail) (productdetailOut PDmodel.ProductDetail, err error)
	Update(ctx context.Context, db *gorm.DB, productdetail PDmodel.ProductDetail) (err error)
	UpdateStock(ctx context.Context, db *gorm.DB, ID uint, operation string, quantity int) (err error)
	Delete(ctx context.Context, db *gorm.DB, ID uint) (err error)
}
