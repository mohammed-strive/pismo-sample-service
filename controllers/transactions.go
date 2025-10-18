package controllers

import (
	"fmt"
	"net/http"
	"pismo-service/models"
	"pismo-service/services"

	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	CreateTransaction(*gin.Context)
}

type transactionController struct {
	service services.TransactionService
}

func NewTransactionController(tcs services.TransactionService) TransactionController {
	return transactionController{
		service: tcs,
	}
}

func (tcc transactionController) CreateTransaction(c *gin.Context) {
	var transationRequestBody models.CreateTransactionJSONBody

	if err := c.ShouldBindBodyWithJSON(&transationRequestBody); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTransaction := models.Transaction{
		AccountId:       transationRequestBody.AccountId,
		OperationTypeId: transationRequestBody.OperationTypeId,
		Amount:          transationRequestBody.Amount,
	}

	result, err := tcc.service.CreateTransaction(c.Request.Context(), newTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
