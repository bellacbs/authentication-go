package user_service

import (
	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain user_model.UserDomainInterface) *rest_errors.RestError {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	err := userDomain.EncryptPassword()
	if err != nil {
		return rest_errors.NewInternalServerError("Error to hash password")
	}
	return nil
}
