package product

import (
	"context"

	MDproduct "github.com/arudji1211/ecommerce/module/model/product"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll(ctx context.Context, db *gorm.DB) (productOut []MDproduct.Product, err error)
	GetByID(ctx context.Context, db *gorm.DB, ID uint) (productOut MDproduct.Product, err error)
	Create(ctx context.Context, db *gorm.DB, productIn MDproduct.Product) (MDproduct.Product, error)
	Update(ctx context.Context, db *gorm.DB, productIn MDproduct.Product) (err error)
	Delete(ctx context.Context, db *gorm.DB, ID uint) (err error)
	IsTransaction(Db *gorm.DB)
}
