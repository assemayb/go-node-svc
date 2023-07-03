package main

import (
	"context"
	"go-svc-test/controllers"
	"go-svc-test/database"
	"go-svc-test/services"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
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
	cpuUsageGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_svc_cpu_usage",
		Help: "CPU usage for Go service",
	})

	memoryUsageGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_svc_memory_usage",
		Help: "Memory usage for Go service",
	})
	prometheus.MustRegister(cpuUsageGauge)
	prometheus.MustRegister(memoryUsageGauge)

	server.GET("/metrics", func(ctx *gin.Context) {

		cpuUsage, err := cpu.Percent(0, false)
		if err != nil {
			log.Fatal(err)
		}
		cpuUsageGauge.Set(cpuUsage[0])

		memoryInfo, err := mem.VirtualMemory()
		if err != nil {
			log.Fatal(err)
		}
		memoryUsage := float64(memoryInfo.Used) / 1024 / 1024
		memoryUsageGauge.Set(memoryUsage)

		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)

	})

	defer mongoClient.Disconnect(ctx)
	basePath := server.Group("/v1")
	transactionsController.RegisterTransactionsRoutes(basePath)
	log.Fatal(server.Run(":3002"))
}
