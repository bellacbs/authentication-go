package user_service

import (
	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByID(userId string) (user_model.UserDomainInterface, *rest_errors.RestError) {
	logger.Info("Init findUserByID services", zap.String("journey", "findUserById"))
	user, err := ud.userRepository.FindUserByID(userId)
	return user, err
}

func (ud *userDomainService) FindUserByEmail(email string) (user_model.UserDomainInterface, *rest_errors.RestError) {
	logger.Info("Init findUserByEmail services", zap.String("journey", "findUserByEmail"))
	user, err := ud.userRepository.FindUserByEmail(email)
	return user, err
}
