package userController

import (
	"net/http"

	"github.com/bellacbs/authentication-go/src/configuration/validation"
	"github.com/bellacbs/authentication-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)

		c.JSON(restError.Code, restError)
	}

	c.JSON(http.StatusOK, userRequest)

}
