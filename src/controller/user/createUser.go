package userController

import (
	"fmt"

	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"github.com/bellacbs/authentication-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := rest_errors.NewBadRequestError(fmt.Sprintf("There  are some incorrect fields, error=%s\n", err.Error()))

		c.JSON(restError.Code, restError)
	}

	fmt.Println(userRequest)

}
