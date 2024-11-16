package product

import (
	"context"

	PDmodel "github.com/arudji1211/ecommerce/module/model/product"
)

type ProductService interface {
	GetAll(ctx context.Context) (Products []PDmodel.Product, err error)
	GetByID(ctx context.Context, ID uint) (Product PDmodel.Product, err error)
	Create(ctx context.Context, ProductIn PDmodel.Product) (ProductOut PDmodel.Product, err error)
	Update(ctx context.Context, Product PDmodel.Product) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
