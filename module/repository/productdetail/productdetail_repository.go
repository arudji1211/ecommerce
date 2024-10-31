package productdetail

import (
	"context"

	productdetail "github.com/arudji1211/ecommerce/module/model/productDetail"
)

type ProductDetailRepository interface {
	GetAllByProductID(ctx context.Context, ID uint) (productdetailIn []productdetail.ProductDetail, err error)
	GetByID(ctx context.Context, ID uint) (productdetailIn productdetail.ProductDetail, err error)
	Create(ctx context.Context, productdetailIn productdetail.ProductDetail) (err error)
	Update(ctx context.Context, productdetailIn productdetail.ProductDetail) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
