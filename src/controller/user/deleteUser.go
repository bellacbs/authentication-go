package user_controller

import (
	"net/http"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller", zap.String("journey", "deleteUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to convert Hex to ObjectID",
			err,
			zap.String("journey", "deleteUser Controller"),
			zap.String("useriId", userId),
		)
		errRespose := rest_errors.NewBadRequestError("UserID is not a valid hex")
		c.JSON(errRespose.Code, errRespose)
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	logger.Info("Delete  User succefully", zap.String("journey", "deleteUser"))
	c.Status(http.StatusOK)

}
