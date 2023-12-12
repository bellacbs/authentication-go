package user_model

import (
	"os"
	"strconv"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var (
	cost = "COST"
)

func NewUserDomain(email, password, name string) UserDomainInterface {
	return &UserDomain{
		email, password, name,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
}

func (ud *UserDomain) EncryptPassword() error {
	cost, err := strconv.Atoi(os.Getenv(cost))
	if err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "getEnv"))
		return err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(ud.Password), cost)
	if err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "encryptPassword"))
		return err
	}
	ud.Password = string(bytes)
	return nil
}

type UserDomainInterface interface {
	CreateUser() *rest_errors.RestError
	UpdateUser(string) *rest_errors.RestError
	FindUser(string) (*UserDomain, *rest_errors.RestError)
	DeleteUser(string) *rest_errors.RestError
}
