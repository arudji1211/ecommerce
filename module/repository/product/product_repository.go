package product

import (
	"context"

	MDproduct "github.com/arudji1211/ecommerce/module/model/product"
)

type ProductRepository interface {
	GetAll(ctx context.Context) (productOut []MDproduct.Product, err error)
	GetByID(ctx context.Context, ID uint) (productOut MDproduct.Product, err error)
	Create(ctx context.Context, productIn MDproduct.Product) (MDproduct.Product, error)
	Update(ctx context.Context, productIn MDproduct.Product) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
