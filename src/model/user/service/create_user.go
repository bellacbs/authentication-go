package user_service

import (
	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain user_model.UserDomainInterface) (user_model.UserDomainInterface, string, *rest_errors.RestError) {
	logger.Info("Init createUser service", zap.String("journey", "createUser"))
	user, _ := ud.FindUserByEmail(userDomain.GetEmail())
	if user != nil {
		return nil, "", rest_errors.NewBadRequestError("Email is already registered in another account")
	}
	errHash := userDomain.EncryptPassword()
	if errHash != nil {
		return nil, "", rest_errors.NewInternalServerError("Error to hash password")
	}
	token, err := userDomain.GenerateToken()
	if err != nil {
		return nil, "", err
	}
	userDomainepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, "", rest_errors.NewInternalServerError("Error to create user on database")
	}
	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomain.GetId()),
		zap.String("journey", "createUserService"))
	return userDomainepository, token, nil
}
