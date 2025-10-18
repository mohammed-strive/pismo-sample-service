package routes

import (
	"pismo-service/controllers"

	"github.com/gin-gonic/gin"
)

func AddV1Routes(r *gin.RouterGroup, acc controllers.AccountController, tcc controllers.TransactionController) {
	r.GET("accounts/:accountId", acc.GetAccountByAccountId)
	r.POST("accounts", acc.CreateAccount)
	r.POST("transactions", tcc.CreateTransaction)
}
