package services

import (
	"go-svc-test/models"
	"net/url"
)

type TransactionsFilter map[string]string

type TransactionService interface {
	GetAllTransactions(filter TransactionsFilter, params url.Values) ([]*models.Transaction, error)
	GetTransactionById(id string) (*models.Transaction, error)
}
