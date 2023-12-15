package user_model

import (
	"os"
	"strconv"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (ud *userDomain) EncryptPassword() error {
	cost, err := strconv.Atoi(os.Getenv(cost))
	if err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "getEnv"))
		return err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(ud.password), cost)
	if err != nil {
		logger.Error("Error trying to hashPassword", err,
			zap.String("journey", "encryptPassword"))
		return err
	}
	ud.password = string(bytes)
	return nil
}

func (ud *userDomain) CheckPasswordHash(hashPassword string, password string) *rest_errors.RestError {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		logger.Error("Error trying to check password", err,
			zap.String("journey", "CheckPasswordHash"))
		return rest_errors.NewForbiddenError("Invalid password")
	}
	return nil
}
