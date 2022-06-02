package carts

import (
	"e-shop/src/products"
	"e-shop/src/users"
)

type Cart struct {
	ID         int
	UserID     int
	CartDetail []CartDetail
	User       users.User
}

type CartDetail struct {
	ID        int
	CartID    int
	ProductID int
	Quantity  int
	Product   products.Product
}
