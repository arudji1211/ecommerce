package product

import (
	"context"

	MDproduct "github.com/arudji1211/ecommerce/module/model/product"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func (p *ProductRepositoryImpl) GetAll(ctx context.Context, db *gorm.DB) (productOut []MDproduct.Product, err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.Db.Model(MDproduct.Product{}).Find(productOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductRepositoryImpl) GetByID(ctx context.Context, db *gorm.DB, ID uint) (productOut MDproduct.Product, err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.Db.Model(MDproduct.Product{}).Where("id = ?", ID).Find(productOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductRepositoryImpl) Create(ctx context.Context, db *gorm.DB, productIn MDproduct.Product) (MDproduct.Product, error) {
	//panic("not implemented") // TODO: Implement
	tx := p.Db.Model(MDproduct.Product{}).Clauses(clause.Returning{}).Create(productIn)
	if err := tx.Error; err != nil {
		return productIn, err
	}
	return productIn, nil
}

func (p *ProductRepositoryImpl) Update(ctx context.Context, db *gorm.DB, productIn MDproduct.Product) (err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.Db.Model(MDproduct.Product{}).Where("id = ?", productIn.ID).Updates(&productIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, ID uint) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.Db.Model(MDproduct.Product{}).Where("id = ?", ID).Delete(MDproduct.Product{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductRepositoryImpl) IsTransaction(Db *gorm.DB) {
	if Db != nil {
		p.Db = Db
	}
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		Db: DB,
	}
}
