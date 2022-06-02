package transactions

import (
	"e-shop/src/products"
	"time"
)

type TransactionFormatter struct {
	ID                 int                          `json:"id"`
	TotalPrice         int                          `json:"total_price"`
	PaymentURL         string                       `json:"payment_url"`
	Status             string                       `json:"status"`
	Deadline           time.Time                    `json:"deadline"`
	CreatedAt          time.Time                    `json:"created_at"`
	UpdatedAt          time.Time                    `json:"updated_at"`
	TransactionDetails []TransactionDetailFormatter `json:"transaction_details"`
}

type TransactionDetailFormatter struct {
	ID       int                       `json:"id"`
	Quantity int                       `json:"quantity"`
	Product  products.ProductFormatter `json:"product"`
}

func FormatTransactionDetail(detail TransactionDetail) TransactionDetailFormatter {
	formatter := TransactionDetailFormatter{}
	formatter.ID = detail.ID
	formatter.Quantity = detail.Quantity
	formatter.Product = products.FormatProduct(detail.Product)

	return formatter
}

func FormatTransactionDetails(details []TransactionDetail) []TransactionDetailFormatter {
	var formatter []TransactionDetailFormatter

	for _, detail := range details {
		format := FormatTransactionDetail(detail)
		formatter = append(formatter, format)
	}

	return formatter
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.TotalPrice = transaction.TotalPrice
	formatter.PaymentURL = transaction.PaymentURL
	formatter.Status = transaction.Status
	formatter.Deadline = transaction.Deadline
	formatter.CreatedAt = transaction.CreatedAt
	formatter.UpdatedAt = transaction.UpdatedAt
	formatter.TransactionDetails = FormatTransactionDetails(transaction.TransactionDetails)

	return formatter
}

func FormatTransactions(transactions []Transaction) []TransactionFormatter {
	var formatter []TransactionFormatter

	for _, transaction := range transactions {
		format := FormatTransaction(transaction)
		formatter = append(formatter, format)
	}

	return formatter
}
