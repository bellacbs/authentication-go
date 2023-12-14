package user_controller

import (
	"net/http"
	"net/mail"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"github.com/bellacbs/authentication-go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init findUserById Controller", zap.String("journey", "findUserById"))
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error to validate userId", err, zap.String("journey", "findUserByID"))
		errorMessage := rest_errors.NewBadRequestError("UserID is  not valid")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByID(userId)
	if err != nil {
		logger.Error("Error to get UserDomain on service", err, zap.String("journey", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertDomainToResponse(userDomain)

	logger.Info("FindUserByID succefully", zap.String("journey", "findUserByID"))

	c.JSON(http.StatusOK, response)
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail Controller", zap.String("journey", "findUserByEmail"))
	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error to validate userEmail", err, zap.String("journey", "findUserByEmail"))
		errorMessage := rest_errors.NewBadRequestError("Email is  not valid")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(userEmail)
	if err != nil {
		logger.Error("Error to get UserDomain on service", err, zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertDomainToResponse(userDomain)

	logger.Info("FindUserByID succefully", zap.String("journey", "findUserByID"))

	c.JSON(http.StatusOK, response)

}
