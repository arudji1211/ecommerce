package productdetail

import (
	"context"

	PDmodel "github.com/arudji1211/ecommerce/module/model/productdetail"
)

type ProductDetailService interface {
	GetAll(ctx context.Context, ID uint) (productdetails []PDmodel.ProductDetail, err error)
	GetByID(ctx context.Context, ID uint) (productdetail PDmodel.ProductDetail, err error)
	Create(ctx context.Context, productdetailIn PDmodel.ProductDetail) (productdetailOut PDmodel.ProductDetail, err error)
	Update(ctx context.Context, productdetail PDmodel.ProductDetail) (err error)
	Delete(ctx context.Context, ID uint) (err error)
}
