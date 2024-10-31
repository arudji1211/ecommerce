package productdetail

import (
	"context"

	productdetail "github.com/arudji1211/ecommerce/module/model/productDetail"
	"gorm.io/gorm"
)

type ProductDetailRepositoryImpl struct {
	db *gorm.DB
}

func (p *ProductDetailRepositoryImpl) GetAllByProductID(ctx context.Context, ID uint) (productdetailOut []productdetail.ProductDetail, err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.db.Model(productdetail.ProductDetail{}).Where("ProductID = ?", ID).Find(&productdetailOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductDetailRepositoryImpl) GetByID(ctx context.Context, ID uint) (productdetailOut productdetail.ProductDetail, err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.db.Model(productdetail.ProductDetail{}).Where("ID = ?", ID).Find(&productdetailOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductDetailRepositoryImpl) Create(ctx context.Context, productdetailIn productdetail.ProductDetail) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(productdetail.ProductDetail{}).Create(productdetailIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductDetailRepositoryImpl) Update(ctx context.Context, productdetailIn productdetail.ProductDetail) (err error) {
	//panic("not implemented") // TODO: Implement
	tx := p.db.Model(productdetail.ProductDetail{}).Where("ID = ?", productdetailIn.ID).Updates(productdetailIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p *ProductDetailRepositoryImpl) Delete(ctx context.Context, ID uint) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(productdetail.ProductDetail{}).Where("ID = ?", ID).Delete(productdetail.ProductDetail{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func NewProductDetailRepository(Db *gorm.DB) ProductDetailRepository {
	return &ProductDetailRepositoryImpl{
		db: Db,
	}
}
