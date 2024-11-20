package cart

import (
	"context"

	Mcart "github.com/arudji1211/ecommerce/module/model/cart"
	"gorm.io/gorm"

	Rcart "github.com/arudji1211/ecommerce/module/repository/cart"
)

type CartServiceImpl struct {
	CartRepo Rcart.CartRepository
}

func (c *CartServiceImpl) GetAll(ctx context.Context, db *gorm.DB, id uint) (cartOut []Mcart.Cart, err error) {
	//panic("not implemented") // TODO: Implement
	cartOut, err = c.CartRepo.GetAll(ctx, db, id)
	if err != nil {
		return
	}
	return
}

func (c *CartServiceImpl) Create(ctx context.Context, db *gorm.DB, cartIn Mcart.Cart) (Mcart.Cart, error) {
	//panic("not implemented") // TODO: Implement
	cartIn, err := c.CartRepo.Create(ctx, db, cartIn)
	if err != nil {
		return cartIn, err
	}
	return cartIn, err
}

func (c *CartServiceImpl) Update(ctx context.Context, db *gorm.DB, cartIn Mcart.Cart) (err error) {
	//panic("not implemented") // TODO: Implement
	err = c.CartRepo.Update(ctx, db, cartIn)
	if err != nil {
		return
	}
	return
}

func (c *CartServiceImpl) Delete(ctx context.Context, db *gorm.DB, id uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = c.CartRepo.Delete(ctx, db, id)
	if err != nil {
		return
	}
	return
}

func NewCartService(repo Rcart.CartRepository) CartService {
	return &CartServiceImpl{
		CartRepo: repo,
	}
}
