package carts

import (
	"e-shop/src/products"
	"e-shop/src/users"
)

type CartFormatter struct {
	ID         int                   `json:"id"`
	UserID     int                   `json:"user_id"`
	User       users.UserFormatter   `json:"user"`
	CartDetail []CartDetailFormatter `json:"cart_details"`
}

type CartDetailFormatter struct {
	ID       int                       `json:"id"`
	CartID   int                       `json:"cart_id"`
	Quantity int                       `json:"quantity"`
	Product  products.ProductFormatter `json:"product"`
}

func FormatCart(cart Cart) CartFormatter {
	formatter := CartFormatter{}
	formatter.ID = cart.ID
	formatter.UserID = cart.UserID
	formatter.User = users.FormatUser(cart.User, "")
	formatter.CartDetail = FormatCartDetails(cart.CartDetail)

	return formatter
}

func FormatCartDetail(detail CartDetail) CartDetailFormatter {
	formatter := CartDetailFormatter{}
	formatter.ID = detail.ID
	formatter.CartID = detail.CartID
	formatter.Quantity = detail.Quantity
	formatter.Product = products.FormatProduct(detail.Product)

	return formatter
}

func FormatCartDetails(details []CartDetail) []CartDetailFormatter {
	var formatter []CartDetailFormatter

	for _, detail := range details {
		format := FormatCartDetail(detail)
		formatter = append(formatter, format)
	}

	return formatter
}
