package models

// Account defines model for Account.
type Account struct {
	AccountId      uint64 `gorm:"primaryKey;column:account_id" json:"account_id"`
	DocumentNumber string `gorm:"unique;column:document_number;not null" json:"document_number"`
}

// CreateAccountJSONBody defines parameters for CreateAccount.
type CreateAccountJSONBody struct {
	DocumentNumber string `json:"document_number"`
}

// CreateAccountJSONRequestBody defines body for CreateAccount for application/json ContentType.
type CreateAccountJSONRequestBody CreateAccountJSONBody
