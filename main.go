package main

import (
	"context"
	"os"

	"github.com/bellacbs/authentication-go/src/configuration/database/mongodb"
	"github.com/bellacbs/authentication-go/src/configuration/logger"
	"github.com/bellacbs/authentication-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.FatalError("Error loading .env file", err)
	}
	port := os.Getenv("PORT")
	dataBase, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.FatalError("Error to connect database", err)
	}
	userController := initDependencies(dataBase)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(port); err != nil {
		logger.FatalError("Error to initialize routes", err)
	}
}
