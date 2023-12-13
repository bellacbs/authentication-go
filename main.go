package main

import (
	"context"
	"os"

	"github.com/bellacbs/authentication-go/src/configuration/database/mongodb"
	"github.com/bellacbs/authentication-go/src/configuration/logger"
	"github.com/bellacbs/authentication-go/src/controller/routes"
	user_controller "github.com/bellacbs/authentication-go/src/controller/user"
	user_service "github.com/bellacbs/authentication-go/src/model/user/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.FatalError("Error loading .env file", err)
	}
	port := os.Getenv("PORT")
	_, err = mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.FatalError("error to connect database", err)
	}
	userService := user_service.NewUserDomainService()
	userController := user_controller.NewUserControllerInterace(userService)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(port); err != nil {
		logger.FatalError("Error to initialize routes", err)
	}
}
