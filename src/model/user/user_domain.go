package user_model

import (
	"os"
	"strconv"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var (
	cost = "COST"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string

	EncryptPassword() error
}

func NewUserDomain(email, password, name string) UserDomainInterface {
	return &userDomain{
		email, password, name,
	}
}

type userDomain struct {
	email    string
	password string
	name     string
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}

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
