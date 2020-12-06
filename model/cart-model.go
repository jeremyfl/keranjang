package model

// Cart is main model of cart
type Cart struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
}
