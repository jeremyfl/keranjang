package controller

import (
	"github.com/jeremylombogia/keranjang/internal"
	"github.com/jeremylombogia/keranjang/request"
	"github.com/labstack/echo/v4"
)

// NewCartController new controller entities
func NewCartController(service internal.ICartService) *CartController {
	return &CartController{
		Service: service,
	}
}

// CartController new entities of cart controller
type CartController struct {
	Service internal.ICartService
}

// Get request controller get
func (cc CartController) Get(c echo.Context) error {
	cart, err := cc.Service.GetCart()
	if err != nil {
		return GenerateResponse(c, 500, err.Error(), nil)
	}

	return GenerateResponse(c, 200, "", cart)
}

// Post request controller post
func (cc CartController) Post(c echo.Context) error {
	request := new(request.CartRequest)
	if err := c.Bind(request); err != nil {
		return GenerateResponse(c, 400, "something went wrong with your request", request)
	}

	cart, err := cc.Service.StoreCart(request)
	if err != nil {
		return GenerateResponse(c, 500, "something went wrong with our system", nil)
	}

	return GenerateResponse(c, 201, "cart created", cart)
}
