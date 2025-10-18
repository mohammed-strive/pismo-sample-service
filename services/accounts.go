package services

import (
	"context"
	"fmt"
	"pismo-service/models"
	"pismo-service/repository"
)

//go:generate mockgen -source=accounts.go -destination=../mocks/mock_accounts_service.go -package=mocks

type AccountService interface {
	GetAccountById(context.Context, uint64) (models.Account, error)
	CreateAccount(context.Context, models.Account) (models.Account, error)
}

func NewAccountService(acr repository.AccountRepository) AccountService {
	return accountService{
		repo: acr,
	}
}

type accountService struct {
	repo repository.AccountRepository
}

func (accs accountService) GetAccountById(c context.Context, accountId uint64) (models.Account, error) {
	result, err := accs.repo.GetOneAccountById(c, accountId)
	if err != nil {
		return result, fmt.Errorf("Account Service: GetAccountById: Error getting Account with ID-%d: %v", accountId, err)
	}

	return result, nil
}

func (accs accountService) CreateAccount(c context.Context, newAccount models.Account) (models.Account, error) {
	result, err := accs.repo.CreateNewAccount(c, newAccount)
	if err != nil {
		return result, fmt.Errorf("Account Service: CreateAccount: Error creating new account: %v", err)
	}

	return result, nil
}
