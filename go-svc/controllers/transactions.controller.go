package controllers

import (
	"go-svc-test/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) TransactionController {
	return TransactionController{
		TransactionService: transactionService,
	}
}

func (tc *TransactionController) GetAllTransactions(ctx *gin.Context) {
	var body map[string]string
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	params := ctx.Request.URL.Query()
	transactions, err := tc.TransactionService.GetAllTransactions(body, params)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Error while fetching transactions"})
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) GetTransactionById(ctx *gin.Context) {
	var id string = ctx.Param("id")
	transaction, err := tc.TransactionService.GetTransactionById(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}

func (tc *TransactionController) RegisterTransactionsRoutes(router *gin.RouterGroup) {
	txnsRouter := router.Group("/transactions")
	txnsRouter.POST("/", tc.GetAllTransactions)
	txnsRouter.GET("/:id", tc.GetTransactionById)
}