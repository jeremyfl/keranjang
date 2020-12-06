package repositories

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/jeremylombogia/keranjang/model"
)

// NewCartRepo initialize cart repositories
func NewCartRepo(dbClient *redis.Client) *CartRepoModel {
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

	var cart model.Cart
	if err := json.Unmarshal([]byte(result), &cart); err != nil {
		panic(err.Error())
	}

	return &[]model.Cart{
		cart,
	}, nil
}

// Insert inserting the repo to database
func (cr CartRepoModel) Insert(payload *model.Cart) (*model.Cart, error) {
	value, _ := json.Marshal(payload)

	cr.Client.Set(ctx, "cart", value, 0).Err()

	return payload, nil
}
