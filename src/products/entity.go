package products

import (
	"e-shop/src/merchants"
	"time"
)

type Product struct {
	ID              int
	MerchantsID     int
	CategoryID      int
	Name            string
	Price           int
	DiscountedPrice int
	Description     string
	Stocks          int
	Merchants       merchants.Merchant
	ProductImages   []ProductImage
	Category        Category
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ProductImage struct {
	ID        int
	ImageURL  string
	ProductID int
}
type Category struct {
	ID          int
	Name        string
	Description string
}
