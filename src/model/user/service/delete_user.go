package user_service

import (
	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_errors.RestError {
	logger.Info("Init DeleteUser service", zap.String("journey", "deleteUserService"))
	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		return err
	}
	logger.Info(
		"DeleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUserService"))
	return nil
}
