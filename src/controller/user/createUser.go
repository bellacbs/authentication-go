package userController

import (
	"net/http"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	"github.com/bellacbs/authentication-go/src/configuration/validation"
	"github.com/bellacbs/authentication-go/src/controller/model/request"
	"github.com/bellacbs/authentication-go/src/controller/model/response"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	UserDomainInterface user_model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restError := validation.ValidateUserError(err)

		c.JSON(restError.Code, restError)
		return
	}

	domain := user_model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := response.UserResponse{
		ID:    uuid.New().String(),
		Email: userRequest.Email,
		Name:  userRequest.Name,
	}

	logger.Info("User created succefully", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)

}
