package user_controller

import (
	"net/http"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"github.com/bellacbs/authentication-go/src/configuration/validation"
	"github.com/bellacbs/authentication-go/src/controller/model/request"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"github.com/bellacbs/authentication-go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller", zap.String("journey", "updateUser"))
	var userRequest request.UserUpdateRequest

	userId := c.Param("userId")
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		restError := validation.ValidateUserError(err)

		c.JSON(restError.Code, restError)
		return
	}

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to convert Hex to ObjectID",
			err,
			zap.String("journey", "updateUser Controller"),
			zap.String("useriId", userId),
		)
		errRespose := rest_errors.NewBadRequestError("UserID is not a valid hex")
		c.JSON(errRespose.Code, errRespose)
	}

	domain := user_model.NewUserUpdateDomain(
		userRequest.Name,
	)
	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	userDomain, err := uc.service.FindUserByID(userId)
	if err != nil {
		logger.Error("Error to get UserDomain on service", err, zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertDomainToResponse(userDomain)
	logger.Info("User created succefully", zap.String("journey", "updateUser"))
	c.JSON(http.StatusOK, response)

}
