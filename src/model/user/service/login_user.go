package user_service

import (
	"fmt"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginService(userDomain user_model.UserDomainInterface) (user_model.UserDomainInterface, string, *rest_errors.RestError) {
	logger.Info("Init LoginUser service", zap.String("journey", "LoginUser"))
	user, _ := ud.FindUserByEmail(userDomain.GetEmail())
	if user == nil {
		return nil, "", rest_errors.NewBadRequestError(fmt.Sprintf("Email %s doesn't exist", userDomain.GetEmail()))
	}
	err := userDomain.CheckPasswordHash(user.GetPassword(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}
	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}
	logger.Info(
		"LoginUser service executed successfully",
		zap.String("userId", userDomain.GetId()),
		zap.String("journey", "loginUserService"))
	return user, token, nil
}
