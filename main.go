package main

import (
	"fmt"
	"os"
	"pismo-service/controllers"
	"pismo-service/middlewares"
	"pismo-service/repository"
	"pismo-service/routes"
	"pismo-service/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middlewares.LoggingMiddleware(logger))

	// setup the database.
	db, err := getDBConnection()
	if err != nil {
		logrus.Fatalf("Error connecting to DB: %v", err)
	}

	// Instantiate repository layer
	accountRepository := repository.NewAccountRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	// Instantiate services
	accountService := services.NewAccountService(accountRepository)
	transactionService := services.NewTransactionService(transactionRepository)

	// Instantiate controllers
	accountController := controllers.NewAccountController(accountService)
	transactionController := controllers.NewTransactionController(transactionService)

	v1Group := router.Group("/v1")

	// Add routes
	routes.AddV1Routes(v1Group, accountController, transactionController)

	noRouteController := controllers.NewNoRouteController()
	router.NoRoute(noRouteController.NoRouteHandler)

	_ = router.Run(":8082")
}

func getDBConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("DB connection successful")
	return db, nil
}
