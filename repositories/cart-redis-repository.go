package repositories

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/jeremylombogia/keranjang/model"
)

// NewCartRedisRepo initialize cart repositories
func NewCartRedisRepo(dbClient *redis.Client) *CartRepoModel {
	return &CartRepoModel{dbClient}
}

// CartRepoModel is cart repo model
type CartRepoModel struct {
	Client *redis.Client
}

var ctx = context.Background()

// Get listing the repo from database
func (cr CartRepoModel) Get() (*[]model.Cart, error) {
	result, err := cr.Client.Get(ctx, "cart").Result()
	if err != nil {
		panic(err.Error())
	}

	var cart []model.Cart
	if err := json.Unmarshal([]byte(result), &cart); err != nil {
		panic(err.Error())
	}

	return &cart, nil
}

// Insert inserting the repo to database
func (cr CartRepoModel) Insert(payload *model.Cart) (*model.Cart, error) {
	var newCart []model.Cart

	previousCart, err := cr.Get()
	if err != nil {
		return nil, err
	}

	for _, carts := range *previousCart {
		newCart = append(newCart, carts)
		newCart = append(newCart, *payload)
	}

	value, err := json.Marshal(newCart)
	if err != nil {
		return nil, err
	}

	if err = cr.Client.Set(ctx, "cart", value, 0).Err(); err != nil {
		return nil, err
	}

	return payload, nil
}
