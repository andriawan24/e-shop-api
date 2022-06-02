package transactions

import (
	"e-shop/src/payment"
	"e-shop/src/users"
	"errors"
	"strconv"
)

type Service interface {
	GetTransactions(userID int) ([]Transaction, error)
	CheckoutCart(transaction Transaction, transactionDetail []TransactionDetail, user users.User) (Transaction, error)
	ProccessPayment(input TransactionNotificationInput) error
}

type service struct {
	repository     Repository
	paymentService payment.Service
}

func NewService(repository Repository, paymentService payment.Service) *service {
	return &service{repository, paymentService}
}

func (s *service) GetTransactions(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetUserTransactions(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CheckoutCart(transaction Transaction, transactionDetail []TransactionDetail, user users.User) (Transaction, error) {

	newTransaction, err := s.repository.InsertTransaction(transaction)
	if err != nil {
		return newTransaction, err
	}

	inputTransaction := payment.Transaction{
		Amount: newTransaction.TotalPrice,
		ID:     newTransaction.ID,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(inputTransaction, user)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	var newDetails []TransactionDetail
	for _, detail := range transactionDetail {
		detail.TransactionID = newTransaction.ID
		newDetails = append(newDetails, detail)
	}

	success, err := s.repository.InsertTransactionDetails(newDetails)
	if err != nil {
		return newTransaction, err
	}

	if len(success) == 0 {
		return newTransaction, errors.New("Cannot insert transaction")
	}

	newTransaction.TransactionDetails = success
	updateTransaction := newTransaction

	return updateTransaction, nil
}

func (s *service) ProccessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByID(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "CANCELLED"
	}

	_, err = s.repository.Update(transaction)
	if err != nil {
		return err
	}

	return nil
}
