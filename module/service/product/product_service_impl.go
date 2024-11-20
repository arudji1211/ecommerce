package product

import (
	"context"

	PDmodel "github.com/arudji1211/ecommerce/module/model/product"
	PRepo "github.com/arudji1211/ecommerce/module/repository/product"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepo PRepo.ProductRepository
}

func (p *ProductServiceImpl) GetAll(ctx context.Context, db *gorm.DB) (Products []PDmodel.Product, err error) {
	//panic("not implemented") // TODO: Implement
	Products, err = p.ProductRepo.GetAll(ctx, db)
	if err != nil {
		return
	}
	return
}

func (p *ProductServiceImpl) GetByID(ctx context.Context, db *gorm.DB, ID uint) (Product PDmodel.Product, err error) {
	// panic("not implemented") // TODO: Implement
	Product, err = p.ProductRepo.GetByID(ctx, db, ID)
	if err != nil {
		return
	}
	return
}

func (p *ProductServiceImpl) Create(ctx context.Context, db *gorm.DB, ProductIn PDmodel.Product) (ProductOut PDmodel.Product, err error) {
	// panic("not implemented") // TODO: Implement
	ProductOut, err = p.ProductRepo.Create(ctx, db, ProductIn)
	if err != nil {
		return
	}
	return
}

func (p *ProductServiceImpl) Update(ctx context.Context, db *gorm.DB, Product PDmodel.Product) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.ProductRepo.Update(ctx, db, Product)
	if err != nil {
		return
	}
	return
}

func (p *ProductServiceImpl) Delete(ctx context.Context, db *gorm.DB, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.ProductRepo.Delete(ctx, db, ID)
	if err != nil {
		return
	}
	return
}

func NewProductService(repo PRepo.ProductRepository) ProductService {
	return &ProductServiceImpl{
		ProductRepo: repo,
	}
}
