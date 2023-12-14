package user_controller

import (
	user_service "github.com/bellacbs/authentication-go/src/model/user/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterace(
	serviceInterface user_service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service user_service.UserDomainService
}
