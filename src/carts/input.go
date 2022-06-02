package carts

type SaveCartInput struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type RemoveProductInput struct {
	ProductID int `json:"product_id" binding:"required"`
}
