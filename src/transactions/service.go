package transactions

import "errors"

type Service interface {
	GetTransactions(userID int) ([]Transaction, error)
	CheckoutCart(transaction Transaction, transactionDetail []TransactionDetail) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactions(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetUserTransactions(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CheckoutCart(transaction Transaction, transactionDetail []TransactionDetail) (Transaction, error) {

	newTransaction, err := s.repository.InsertTransaction(transaction)
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
