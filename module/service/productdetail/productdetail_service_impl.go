package productdetail

import (
	"context"

	PDmodel "github.com/arudji1211/ecommerce/module/model/productdetail"
	PDrepo "github.com/arudji1211/ecommerce/module/repository/productdetail"
)

type ProductDetailServiceImpl struct {
	ProductDetailRepo PDrepo.ProductDetailRepository
}

func (p *ProductDetailServiceImpl) GetAll(ctx context.Context, ID uint) (productdetails []PDmodel.ProductDetail, err error) {
	//panic("not implemented") // TODO: Implement
	productdetails, err = p.ProductDetailRepo.GetAllByProductID(ctx, ID)
	if err != nil {
		return
	}
	return
}

func (p *ProductDetailServiceImpl) GetByID(ctx context.Context, ID uint) (productdetail PDmodel.ProductDetail, err error) {
	//panic("not implemented") // TODO: Implement
	productdetail, err = p.ProductDetailRepo.GetByID(ctx, ID)
	if err != nil {
		return
	}
	return
}

func (p *ProductDetailServiceImpl) Create(ctx context.Context, productdetailIn PDmodel.ProductDetail) (ProductDetailOut PDmodel.ProductDetail, err error) {
	ProductDetailOut, err = p.ProductDetailRepo.Create(ctx, productdetailIn)
	if err != nil {
		return ProductDetailOut, err
	}
	return ProductDetailOut, err
}

func (p *ProductDetailServiceImpl) Update(ctx context.Context, productdetail PDmodel.ProductDetail) (err error) {
	err = p.ProductDetailRepo.Update(ctx, productdetail)
	if err != nil {
		return
	}
	return
}

func (p *ProductDetailServiceImpl) Delete(ctx context.Context, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.ProductDetailRepo.Delete(ctx, ID)
	if err != nil {
		return
	}
	return
}

func NewProductDetailService(repo PDrepo.ProductDetailRepository) ProductDetailService {
	return &ProductDetailServiceImpl{
		ProductDetailRepo: repo,
	}
}
