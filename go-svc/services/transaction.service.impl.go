package services

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"go-svc-test/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionServiceImpl struct {
	transactionsCollection *mongo.Collection
	ctx                    context.Context
}

func NewTransactionService(transactionsCollection *mongo.Collection, ctx context.Context) TransactionService {
	return &TransactionServiceImpl{
		transactionsCollection: transactionsCollection,
		ctx:                    ctx,
	}
}

func (t *TransactionServiceImpl) GetAllTransactions(filter TransactionsFilter, params url.Values) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	limit := int64(15000)

	rps := params.Get("rps")
	rpsInt, err := strconv.ParseInt(rps, 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		limit = rpsInt
		fmt.Println("limit: ", limit)
	}

	var cursor *mongo.Cursor
	if filter == nil {
		searchObject := bson.D{}
		for key, value := range filter {
			searchObject = append(searchObject, bson.E{Key: key, Value: value})
		}
		cursor, err = t.transactionsCollection.Find(t.ctx, searchObject, &options.FindOptions{Limit: &limit})
		if err != nil {
			return nil, err
		}
	} else {
		cursor, err = t.transactionsCollection.Find(t.ctx, bson.D{}, &options.FindOptions{Limit: &limit})
		if err != nil {
			return nil, err
		}
	}

	for cursor.Next(t.ctx) {
		var transaction models.Transaction
		err := cursor.Decode(&transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(t.ctx)

	if len(transactions) == 0 {
		return nil, errors.New("documents not found")
	}
	return transactions, nil

}

func (t *TransactionServiceImpl) GetTransactionById(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	filter := bson.D{{Key: "txnDetailsId", Value: id}}
	err := t.transactionsCollection.FindOne(t.ctx, filter).Decode(&transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
