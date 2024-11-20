package cart

import (
	"context"

	Mcart "github.com/arudji1211/ecommerce/module/model/cart"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func (c *CartRepositoryImpl) GetAll(ctx context.Context, db *gorm.DB, id uint) (cartOut []Mcart.Cart, err error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(Mcart.Cart{}).Find(&cartOut)

	if err = tx.Error; err != nil {
		return
	}
	return
}

func (c *CartRepositoryImpl) Create(ctx context.Context, db *gorm.DB, cartIn Mcart.Cart) (Mcart.Cart, error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(Mcart.Cart{}).Clauses(clause.Returning{}).Create(cartIn)
	if err := tx.Error; err != nil {
		return cartIn, err
	}
	return cartIn, nil
}

func (c *CartRepositoryImpl) Update(ctx context.Context, db *gorm.DB, cartIn Mcart.Cart) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(Mcart.Cart{}).Where("id = ?").Update("quantity", &cartIn.Quantity)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (c *CartRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uint) (err error) {
	tx := c.db.Model(Mcart.Cart{}).Where("id = ?", id).Delete(&Mcart.Cart{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (c *CartRepositoryImpl) IsTransaction(db *gorm.DB) {
	if db != nil {
		c.db = db
	}
}

func NewCartRepository(DB *gorm.DB) CartRepository {
	return &CartRepositoryImpl{
		db: DB,
	}
}
