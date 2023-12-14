package user_model

import (
	"os"
	"strconv"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
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
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "encryptPassword"))
		return err
	}
	ud.password = string(bytes)
	return nil
}
