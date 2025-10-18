package repository

import (
	"context"
	"fmt"
	"pismo-service/models"
	"pismo-service/utils"

	"gorm.io/gorm"
)

//go:generate mockgen -source=accounts.go -destination=../mocks/mock_accounts_repository.go -package=mocks

type AccountRepository interface {
	GetOneAccountById(context.Context, uint64) (models.Account, error)
	CreateNewAccount(context.Context, models.Account) (models.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return accountRepository{
		db,
	}
}

func (acr accountRepository) GetOneAccountById(c context.Context, accountId uint64) (models.Account, error) {
	logger := utils.GetLogger(c)
	logger.Infof("Reading account having ID -> %d", accountId)

	var user models.Account
	result := acr.db.First(&user, accountId)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (acr accountRepository) CreateNewAccount(c context.Context, newAccount models.Account) (models.Account, error) {
	logger := utils.GetLogger(c)
	logger.Infof("Inserting new Account: %v", newAccount)

	account := newAccount
	if err := acr.db.Where(&models.Account{DocumentNumber: account.DocumentNumber}).First(&account).Error; err == nil {
		return account, fmt.Errorf("record already exists")
	}

	result := acr.db.Create(&account)
	if result.Error != nil {
		return account, result.Error
	}

	return account, nil
}
