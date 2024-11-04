package product

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/product"
)

type ProductRepository interface {
	GetAll(ctx context.Context) (productOut []product.Product, err error)
	GetByID(ctx context.Context, ID uint) (productOut product.Product, err error)
	Create(ctx context.Context, productIn product.Product) (product.Product, error)
	Update(ctx context.Context, productIn product.Product) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
