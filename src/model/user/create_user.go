package user_model

import (
	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_errors.RestError {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	err := ud.EncryptPassword()
	if err != nil {
		return rest_errors.NewInternalServerError("Error to hash password")
	}
	return nil
}
