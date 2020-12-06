package internal

import (
	"github.com/jeremylombogia/keranjang/model"
	"github.com/jeremylombogia/keranjang/request"
)

// ICartService invoke cart service
type ICartService interface {
	GetCart() (*[]model.Cart, error)
	StoreCart(payload *request.CartRequest) (*model.Cart, error)
}

// ICartRepo invoke cart repo
type ICartRepo interface {
	Get() (*[]model.Cart, error)
	Insert(payload *model.Cart) (*model.Cart, error)
}
