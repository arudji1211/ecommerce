package productdetail

import (
	"context"

	MDproductdetail "github.com/arudji1211/ecommerce/module/model/productdetail"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductDetailRepositoryImpl struct {
	db *gorm.DB
}

func (p *ProductDetailRepositoryImpl) GetAllByProductID(ctx context.Context, db *gorm.DB, ID uint) (productdetailOut []MDproductdetail.ProductDetail, err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.db.Model(MDproductdetail.ProductDetail{}).Where("ProductID = ?", ID).Find(&productdetailOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductDetailRepositoryImpl) GetByID(ctx context.Context, db *gorm.DB, ID uint) (productdetailOut MDproductdetail.ProductDetail, err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.db.Model(MDproductdetail.ProductDetail{}).Where("ID = ?", ID).Find(&productdetailOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductDetailRepositoryImpl) Create(ctx context.Context, db *gorm.DB, productdetailIn MDproductdetail.ProductDetail) (MDproductdetail.ProductDetail, error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(MDproductdetail.ProductDetail{}).Clauses(clause.Returning{}).Create(productdetailIn)
	if err := tx.Error; err != nil {
		return productdetailIn, err
	}
	return productdetailIn, nil
}

func (p *ProductDetailRepositoryImpl) Update(ctx context.Context, db *gorm.DB, productdetailIn MDproductdetail.ProductDetail) (err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.db.Model(MDproductdetail.ProductDetail{}).Where("ID = ?", productdetailIn.ID).Updates(productdetailIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductDetailRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, ID uint) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(MDproductdetail.ProductDetail{}).Where("ID = ?", ID).Delete(MDproductdetail.ProductDetail{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductDetailRepositoryImpl) IsTransaction(Db *gorm.DB) {
	if Db != nil {
		p.db = Db
	}
}

func NewProductDetailRepository(Db *gorm.DB) ProductDetailRepository {
	return &ProductDetailRepositoryImpl{
		db: Db,
	}
}
