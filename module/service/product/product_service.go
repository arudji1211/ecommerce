package product

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/product"
)

type ProductService interface {
	GetAll(ctx context.Context) (Products []product.Product, err error)
	GetByID(ctx context.Context, ID uint) (Product product.Product, err error)
	Create(ctx context.Context, ProductIn product.Product) (ProductOut product.Product, err error)
	Update(ctx context.Context, Product product.Product) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
