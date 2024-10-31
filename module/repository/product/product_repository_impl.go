package product

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/product"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func (p *ProductRepositoryImpl) GetAll(ctx context.Context) (productOut []product.Product, err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.Db.Model(product.Product{}).Find(productOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductRepositoryImpl) GetByID(ctx context.Context, ID uint) (productOut product.Product, err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.Db.Model(product.Product{}).Where("id = ?", ID).Find(productOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductRepositoryImpl) Create(ctx context.Context, productIn product.Product) (err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.Db.Model(product.Product{}).Create(productIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductRepositoryImpl) Update(ctx context.Context, productIn product.Product) (err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.Db.Model(product.Product{}).Where("id = ?", productIn.ID).Updates(&productIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductRepositoryImpl) Delete(ctx context.Context, ID uint) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.Db.Model(product.Product{}).Where("id = ?", ID).Delete(product.Product{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		Db: DB,
	}
}
