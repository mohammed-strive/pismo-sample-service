package models

// Defines values for OperationType.
const (
	CASH_PURCHASE        OperationType = 1
	INSTALLMENT_PURCHASE OperationType = 2
	WITHDRAWAL           OperationType = 3
	PAYMENT              OperationType = 4
)

type OperationType int

// Transaction defines model for Transaction.
type Transaction struct {
	TransactionId   uint64        `gorm:"primaryKey;column:transaction_id"`
	AccountId       int           `gorm:"column:account_id;not null" json:"account_id"`
	Amount          float32       `gorm:"column:amount;not null" json:"amount"`
	OperationTypeId OperationType `gorm:"not null;column:operationtype_id" json:"operation_type_id"`
}

// CreateTransactionJSONBody defines parameters for CreateTransaction.
type CreateTransactionJSONBody struct {
	AccountId       int           `json:"account_id"`
	Amount          float32       `json:"amount"`
	OperationTypeId OperationType `json:"operation_type_id"`
}

// CreateTransactionJSONRequestBody defines body for CreateTransaction for application/json ContentType.
type CreateTransactionJSONRequestBody CreateTransactionJSONBody
