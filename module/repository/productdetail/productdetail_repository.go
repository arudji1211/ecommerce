package productdetail

import (
	"context"

	MDproductdetail "github.com/arudji1211/ecommerce/module/model/productdetail"
)

type ProductDetailRepository interface {
	GetAllByProductID(ctx context.Context, ID uint) (productdetailIn []MDproductdetail.ProductDetail, err error)
	GetByID(ctx context.Context, ID uint) (productdetailIn MDproductdetail.ProductDetail, err error)
	Create(ctx context.Context, productdetailIn MDproductdetail.ProductDetail) (MDproductdetail.ProductDetail, error)
	Update(ctx context.Context, productdetailIn MDproductdetail.ProductDetail) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
