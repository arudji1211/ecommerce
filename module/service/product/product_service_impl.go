package product

import (
	"context"

	PDmodel "github.com/arudji1211/ecommerce/module/model/product"
	PRepo "github.com/arudji1211/ecommerce/module/repository/product"
)

type ProductServiceImpl struct {
	ProductRepo PRepo.ProductRepository
}

func (p *ProductServiceImpl) GetAll(ctx context.Context) (Products []PDmodel.Product, err error) {
	//panic("not implemented") // TODO: Implement
	Products, err = p.ProductRepo.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

func (p *ProductServiceImpl) GetByID(ctx context.Context, ID uint) (Product PDmodel.Product, err error) {
	// panic("not implemented") // TODO: Implement
	Product, err = p.ProductRepo.GetByID(ctx, ID)
	if err != nil {
		return
	}
	return
}

func (p *ProductServiceImpl) Create(ctx context.Context, ProductIn PDmodel.Product) (ProductOut PDmodel.Product, err error) {
	// panic("not implemented") // TODO: Implement
	ProductOut, err = p.ProductRepo.Create(ctx, ProductIn)
	if err != nil {
		return
	}
	return
}

func (p *ProductServiceImpl) Update(ctx context.Context, Product PDmodel.Product) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.ProductRepo.Update(ctx, Product)
	if err != nil {
		return
	}
	return
}

func (p *ProductServiceImpl) Delete(ctx context.Context, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.ProductRepo.Delete(ctx, ID)
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
