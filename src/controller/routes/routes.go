package routes

import (
	user_controller "github.com/bellacbs/authentication-go/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, userController user_controller.UserControllerInterface) {
	userGroup := router.Group("user")
	userGroup.GET("/getUserById/:userId", userController.FindUserById)
	userGroup.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	userGroup.POST("/", userController.CreateUser)
	userGroup.PUT("/:userId", userController.UpdateUser)
	userGroup.DELETE("/:userId", userController.DeleteUser)
}
