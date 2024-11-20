package productdetail

import (
	"context"

	PDmodel "github.com/arudji1211/ecommerce/module/model/productdetail"
	PDrepo "github.com/arudji1211/ecommerce/module/repository/productdetail"
	"gorm.io/gorm"
)

type ProductDetailServiceImpl struct {
	ProductDetailRepo PDrepo.ProductDetailRepository
}

func (p *ProductDetailServiceImpl) GetAll(ctx context.Context, db *gorm.DB, ID uint) (productdetails []PDmodel.ProductDetail, err error) {
	//panic("not implemented") // TODO: Implement
	productdetails, err = p.ProductDetailRepo.GetAllByProductID(ctx, db, ID)
	if err != nil {
		return
	}
	return
}

func (p *ProductDetailServiceImpl) GetByID(ctx context.Context, db *gorm.DB, ID uint) (productdetail PDmodel.ProductDetail, err error) {
	//panic("not implemented") // TODO: Implement
	productdetail, err = p.ProductDetailRepo.GetByID(ctx, db, ID)
	if err != nil {
		return
	}
	return
}

func (p *ProductDetailServiceImpl) Create(ctx context.Context, db *gorm.DB, productdetailIn PDmodel.ProductDetail) (ProductDetailOut PDmodel.ProductDetail, err error) {
	ProductDetailOut, err = p.ProductDetailRepo.Create(ctx, db, productdetailIn)
	if err != nil {
		return ProductDetailOut, err
	}
	return ProductDetailOut, err
}

func (p *ProductDetailServiceImpl) Update(ctx context.Context, db *gorm.DB, productdetail PDmodel.ProductDetail) (err error) {
	err = p.ProductDetailRepo.Update(ctx, db, productdetail)
	if err != nil {
		return
	}
	return
}

func (p *ProductDetailServiceImpl) UpdateStock(ctx context.Context, db *gorm.DB, ID uint, operation string, quantity int) (err error) {
	err = p.ProductDetailRepo.UpdateStock(ctx, db, ID, operation, quantity)
	if err != nil {
		return
	}
	return
}

func (p *ProductDetailServiceImpl) Delete(ctx context.Context, db *gorm.DB, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.ProductDetailRepo.Delete(ctx, db, ID)
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
