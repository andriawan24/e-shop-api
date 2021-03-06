package products

import "e-shop/src/merchants"

type ProductFormatter struct {
	ID              int                         `json:"id"`
	Name            string                      `json:"name"`
	Price           int                         `json:"price"`
	DiscountedPrice int                         `json:"discounted_price"`
	Description     string                      `json:"description"`
	Stocks          int                         `json:"stocks"`
	Merchant        merchants.MerchantFormatter `json:"merchant"`
	ProductImages   []ProductImageFormatter     `json:"images"`
	ProductCategory CategoryFormatter           `json:"category"`
}

// Format Image Products
type ProductImageFormatter struct {
	ID        int    `json:"id"`
	ImageURL  string `json:"image_url"`
	ProductID int    `json:"products_id"`
}

func formatProductImage(image ProductImage) ProductImageFormatter {
	formatter := ProductImageFormatter{}
	formatter.ID = image.ID
	formatter.ImageURL = image.ImageURL
	formatter.ProductID = image.ProductID
	return formatter
}

func FormatProductImages(images []ProductImage) []ProductImageFormatter {
	var formatterImages []ProductImageFormatter

	for _, image := range images {
		formatterImages = append(formatterImages, formatProductImage(image))
	}

	return formatterImages
}

// Format Category
type ProductCategoryFormatter struct {
	ID       int               `json:"id"`
	Category CategoryFormatter `json:"category"`
}

type CategoryFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func FormatCategory(category Category) CategoryFormatter {
	formatter := CategoryFormatter{}
	formatter.ID = category.ID
	formatter.Name = category.Name
	formatter.Description = category.Description

	return formatter
}

func FormatCategories(categories []Category) []CategoryFormatter {
	var formatter []CategoryFormatter
	for _, category := range categories {
		format := CategoryFormatter{}
		format.ID = category.ID
		format.Name = category.Name
		format.Description = category.Description
		formatter = append(formatter, format)
	}

	return formatter
}

func FormatProduct(product Product) ProductFormatter {
	formatter := ProductFormatter{}
	formatter.ID = product.ID
	formatter.Name = product.Name
	formatter.Price = product.Price
	formatter.DiscountedPrice = product.DiscountedPrice
	formatter.Description = product.Description
	formatter.Stocks = product.Stocks
	formatter.Merchant = merchants.FormatMerchant(product.Merchants)
	formatter.ProductImages = FormatProductImages(product.ProductImages)
	formatter.ProductCategory = FormatCategory(product.Category)

	return formatter
}

func FormatProducts(products []Product) []ProductFormatter {
	var formatter []ProductFormatter
	for _, product := range products {
		formatter = append(formatter, FormatProduct(product))
	}

	return formatter
}
