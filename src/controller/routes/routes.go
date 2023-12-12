package routes

import (
	userController "github.com/bellacbs/authentication-go/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("user")
	userGroup.GET("/getUserById/:userId", userController.FindUserById)
	userGroup.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	userGroup.POST("/", userController.CreateUser)
	userGroup.PUT("/:userId", userController.UpdateUser)
	userGroup.DELETE("/:userId", userController.DeleteUser)
}
