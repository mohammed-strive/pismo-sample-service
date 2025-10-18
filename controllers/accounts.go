package controllers

import (
	"net/http"
	"pismo-service/models"
	"pismo-service/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	GetAccountByAccountId(*gin.Context)
	CreateAccount(*gin.Context)
}

type accountController struct {
	service services.AccountService
}

func NewAccountController(acs services.AccountService) AccountController {
	return accountController{
		service: acs,
	}
}

func (acc accountController) GetAccountByAccountId(c *gin.Context) {
	rawAccountId := c.Params.ByName("accountId")
	if rawAccountId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no accountId provided"})
		return
	}

	accountId, err := strconv.ParseUint(rawAccountId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid accountId"})
		return
	}

	result, err := acc.service.GetAccountById(c.Request.Context(), accountId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
func (acc accountController) CreateAccount(c *gin.Context) {
	var accountRequestBody models.CreateAccountJSONBody
	if err := c.ShouldBindBodyWithJSON(&accountRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAccount := models.Account{
		DocumentNumber: accountRequestBody.DocumentNumber,
	}

	result, err := acc.service.CreateAccount(c.Request.Context(), newAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
