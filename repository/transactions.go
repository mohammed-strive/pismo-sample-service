package repository

import (
	"context"
	"fmt"
	"pismo-service/models"
	"pismo-service/utils"

	"gorm.io/gorm"
)

//go:generate mockgen -source=transactions.go -destination=../mocks/mock_transactions_repository.go -package=mocks

type TransactionRepository interface {
	AddNewTransaction(context.Context, models.Transaction) (models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return transactionRepository{
		db,
	}
}

func (tar transactionRepository) AddNewTransaction(c context.Context, newTransaction models.Transaction) (models.Transaction, error) {
	logger := utils.GetLogger(c)
	logger.Infof("Creating new transaction: %v", newTransaction)

	transaction := newTransaction
	if err := tar.db.Where(&models.Transaction{AccountId: transaction.AccountId, Amount: transaction.Amount, OperationTypeId: transaction.OperationTypeId}).First(&transaction).Error; err == nil {
		fmt.Printf("Record exists: %v", err)
		return transaction, fmt.Errorf("record already exists")
	}
	result := tar.db.Create(&transaction)

	if result.Error != nil {
		return transaction, result.Error
	}

	return transaction, nil
}
