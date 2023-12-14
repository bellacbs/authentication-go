package main

import (
	user_controller "github.com/bellacbs/authentication-go/src/controller/user"
	user_repository "github.com/bellacbs/authentication-go/src/model/user/repository"
	user_service "github.com/bellacbs/authentication-go/src/model/user/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependecies(
	database *mongo.Database,
) user_controller.UserControllerInterface {
	userRepository := user_repository.NewUserRepository(database)
	userService := user_service.NewUserDomainService(userRepository)
	return user_controller.NewUserControllerInterace(userService)
}
