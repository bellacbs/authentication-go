package user_service

import (
	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain user_model.UserDomainInterface) *rest_errors.RestError {
	logger.Info("Init updateUser service", zap.String("journey", "updateUserService"))
	err := ud.userRepository.UpdateUserById(userId, userDomain)
	if err != nil {
		return rest_errors.NewInternalServerError("Error to create user on database")
	}
	logger.Info(
		"Update service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUserService"))
	return nil
}
