package product

import (
	"context"

	PDmodel "github.com/arudji1211/ecommerce/module/model/product"
	"gorm.io/gorm"
)

type ProductService interface {
	GetAll(ctx context.Context, db *gorm.DB) (Products []PDmodel.Product, err error)
	GetByID(ctx context.Context, db *gorm.DB, ID uint) (Product PDmodel.Product, err error)
	Create(ctx context.Context, db *gorm.DB, ProductIn PDmodel.Product) (ProductOut PDmodel.Product, err error)
	Update(ctx context.Context, db *gorm.DB, Product PDmodel.Product) (err error)
	Delete(ctx context.Context, db *gorm.DB, ID uint) (err error)
}
