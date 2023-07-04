package main

import (
	"context"
	"fmt"
	"go-svc-test/controllers"
	"go-svc-test/database"
	"go-svc-test/services"
	"os"
	"time"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/process"
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

	go func() {
		for {
			pid := os.Getpid()
			p, _ := process.NewProcess(int32(pid))

			memInfo, err := p.MemoryInfo()

			if err != nil {
				fmt.Println("Error getting memory info:", err)
				return
			}
			cpuPercent, err := p.CPUPercent()
			if err != nil {
				fmt.Println("Error getting CPU percent:", err)
				return
			}

			cpuUsageGauge.Set(cpuPercent)
			memoryUsageGauge.Set(float64(memInfo.RSS) / 1024 / 1024)
			fmt.Println("CPU Percent:", cpuPercent, "Memory Info:", memInfo.RSS/1024/1024, "MB")
			time.Sleep(1 * time.Second)
		}
	}()

	server.GET("/metrics", func(ctx *gin.Context) {
		// fmt.Println("Metrics endpoint called")
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})

	defer mongoClient.Disconnect(ctx)
	basePath := server.Group("/v1")
	transactionsController.RegisterTransactionsRoutes(basePath)
	log.Fatal(server.Run(":3002"))
}
