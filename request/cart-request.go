package request

// CartRequest is HTTP request with JSON
type CartRequest struct {
	ProductID int    `json:"product_id" form:"product_id" query:"product_id"`
	Name      string `json:"name" form:"name" query:"name"`
	Color     string `json:"color" form:"color" query:"color"`
}
