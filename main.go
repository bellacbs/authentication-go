package main

import (
	"log"
	"os"

	"github.com/bellacbs/authentication-go/src/controller/routes"
	user_controller "github.com/bellacbs/authentication-go/src/controller/user"
	user_service "github.com/bellacbs/authentication-go/src/model/user/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	userService := user_service.NewUserDomainService()
	userController := user_controller.NewUserControllerInterace(userService)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}
