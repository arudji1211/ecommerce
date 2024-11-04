package cart

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/cart"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func (c *CartRepositoryImpl) GetAll(ctx context.Context, id uint) (cartOut []cart.Cart, err error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(cart.Cart{}).Find(&cartOut)

	if err = tx.Error; err != nil {
		return
	}
	return
}

func (c *CartRepositoryImpl) Create(ctx context.Context, cartIn cart.Cart) (cart.Cart, error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(cart.Cart{}).Clauses(clause.Returning{}).Create(cartIn)
	if err := tx.Error; err != nil {
		return cartIn, err
	}
	return cartIn, nil
}

func (c *CartRepositoryImpl) Update(ctx context.Context, cartIn cart.Cart) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(cart.Cart{}).Where("id = ?").Update("quantity", &cartIn.Quantity)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (c *CartRepositoryImpl) Delete(ctx context.Context, id uint) (err error) {
	tx := c.db.Model(cart.Cart{}).Where("id = ?", id).Delete(&cart.Cart{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func NewCartRepository(DB *gorm.DB) CartRepository {
	return &CartRepositoryImpl{
		db: DB,
	}
}
