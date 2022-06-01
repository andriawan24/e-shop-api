package products

import (
	"e-shop/src/merchants"
	"time"
)

type Product struct {
	ID                int
	MerchantsID       int
	Name              string
	Price             int
	DiscountedPrice   int
	Description       string
	Stocks            int
	Merchants         merchants.Merchant
	ProductImages     []ProductImage
	ProductCategories []ProductCategory
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type ProductImage struct {
	ID        int
	ImageURL  string
	ProductID int
}

type ProductCategory struct {
	ID         int
	ProductID  int
	CategoryID int
	Category   Category
}

type Category struct {
	ID          int
	Name        string
	Description string
}
