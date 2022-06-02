package transactions

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	GetUserTransactions(userID int) ([]Transaction, error)
	InsertTransaction(transaction Transaction) (Transaction, error)
	InsertTransactionDetails(details []TransactionDetail) ([]TransactionDetail, error)
	GetByID(ID int) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUserTransactions(userID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("TransactionDetails", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Preload(clause.Associations)
		})
	}).Preload(clause.Associations).Find(&transactions).Error

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) InsertTransaction(transaction Transaction) (Transaction, error) {

	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	err = r.db.Where("id = ?", transaction.ID).Preload("TransactionDetails", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Preload(clause.Associations)
		})
	}).Preload(clause.Associations).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) InsertTransactionDetails(details []TransactionDetail) ([]TransactionDetail, error) {

	var transactionDetails []TransactionDetail
	err := r.db.Create(&details).Error
	if err != nil {
		return transactionDetails, err
	}

	err = r.db.Where("transaction_id = ?", details[0].TransactionID).Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Preload(clause.Associations)
	}).Preload(clause.Associations).Find(&transactionDetails).Error
	if err != nil {
		return transactionDetails, err
	}

	return transactionDetails, nil
}

func (r *repository) GetByID(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Where("id = ?", ID).Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
