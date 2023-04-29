package main

import (
	"context"
	"go-svc-test/controllers"
	"go-svc-test/database"
	"go-svc-test/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server                 *gin.Engine
	transactionsService    services.TransactionService
	transactionsController controllers.TransactionController
	ctx                    context.Context
	transactionsCollection *mongo.Collection
	mongoClient            *mongo.Client
)

func init() {
	mongoClient := database.ConnectDB()
	transactionsCollection = mongoClient.Database("axisPayCore").Collection("transactionDetails")
	transactionsService = services.NewTransactionService(transactionsCollection, ctx)
	transactionsController = controllers.NewTransactionController(transactionsService)

	server = gin.New()
	server.Use(gin.Logger(), gin.Recovery())

}

func main() {

	server.GET("/_health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	defer mongoClient.Disconnect(ctx)
	basePath := server.Group("/v1")
	transactionsController.RegisterTransactionsRoutes(basePath)
	log.Fatal(server.Run(":9090"))
}
