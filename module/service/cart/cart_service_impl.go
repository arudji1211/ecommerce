package cart

import (
	"context"

	Mcart "github.com/arudji1211/ecommerce/module/model/cart"

	Rcart "github.com/arudji1211/ecommerce/module/repository/cart"
)

type CartServiceImpl struct {
	CartRepo Rcart.CartRepository
}

func (c *CartServiceImpl) GetAll(ctx context.Context, id uint) (cartOut []Mcart.Cart, err error) {
	//panic("not implemented") // TODO: Implement
	cartOut, err = c.CartRepo.GetAll(ctx, id)
	if err != nil {
		return
	}
	return
}

func (c *CartServiceImpl) Create(ctx context.Context, cartIn Mcart.Cart) (Mcart.Cart, error) {
	//panic("not implemented") // TODO: Implement
	cartIn, err := c.CartRepo.Create(ctx, cartIn)
	if err != nil {
		return cartIn, err
	}
	return cartIn, err
}

func (c *CartServiceImpl) Update(ctx context.Context, cartIn Mcart.Cart) (err error) {
	//panic("not implemented") // TODO: Implement
	err = c.CartRepo.Update(ctx, cartIn)
	if err != nil {
		return
	}
	return
}

func (c *CartServiceImpl) Delete(ctx context.Context, id uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = c.CartRepo.Delete(ctx, id)
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
