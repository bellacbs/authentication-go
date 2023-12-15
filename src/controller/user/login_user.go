package user_controller

import (
	"net/http"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	"github.com/bellacbs/authentication-go/src/configuration/validation"
	"github.com/bellacbs/authentication-go/src/controller/model/request"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"github.com/bellacbs/authentication-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller", zap.String("journey", "createUser"))
	var userLoginRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "loginUser"))
		restError := validation.ValidateUserError(err)

		c.JSON(restError.Code, restError)
		return
	}

	domain := user_model.NewUserLoginDomain(
		userLoginRequest.Email,
		userLoginRequest.Password,
	)

	domainResult, token, err := uc.service.LoginService(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertDomainToResponse(domainResult)

	logger.Info("Login User succefully", zap.String("journey", "loginUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, response)

}
