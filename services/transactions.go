package services

import (
	"context"
	"fmt"
	"pismo-service/models"
	"pismo-service/repository"
)

//go:generate mockgen -source=transactions.go -destination=../mocks/mock_transaction_service.go -package=mocks

type TransactionService interface {
	CreateTransaction(context.Context, models.Transaction) (models.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return transactionService{
		repo,
	}
}

func (tcs transactionService) CreateTransaction(c context.Context, newTransaction models.Transaction) (models.Transaction, error) {
	result, err := tcs.repo.AddNewTransaction(c, newTransaction)
	if err != nil {
		return result, fmt.Errorf("TransactionService: CreateTransaction: Error creating new Transaction: %v", err)
	}

	return result, nil
}
