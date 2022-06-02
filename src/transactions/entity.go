package transactions

import (
	"e-shop/src/products"
	"time"
)

type Transaction struct {
	ID                 int
	UserID             int
	TotalPrice         int
	PaymentURL         string
	Status             string
	Deadline           time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
	TransactionDetails []TransactionDetail
}

type TransactionDetail struct {
	ID            int
	TransactionID int
	ProductID     int
	Quantity      int
	Product       products.Product
}
