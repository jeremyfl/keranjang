package service

import (
	"github.com/jeremylombogia/keranjang/internal"
	"github.com/jeremylombogia/keranjang/model"
	"github.com/jeremylombogia/keranjang/request"
)

// NewCartService new cart service entities
func NewCartService(repo internal.ICartRepo) *CartService {
	return &CartService{repo}
}

// CartService is service for cart
type CartService struct {
	Repository internal.ICartRepo
}

// GetCart service for listing
func (cs CartService) GetCart() (*[]model.Cart, error) {
	cart, err := cs.Repository.Get()
	if err != nil {
		return cart, err
	}

	return cart, nil
}

// StoreCart service for inserting cart to repo
func (cs CartService) StoreCart(payload *request.CartRequest) *model.Cart {
	value := &model.Cart{
		ProductID: payload.ProductID,
		Name:      payload.Name,
		Color:     payload.Color,
	}

	cs.Repository.Insert(value)

	return value
}
