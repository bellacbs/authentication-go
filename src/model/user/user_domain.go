package user_model

import (
	"encoding/json"
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

	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword() error
}

func NewUserDomain(email, password, name string) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
	}
}

type userDomain struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (ud *userDomain) GetJSONValue() (string, error) {
	value, err := json.Marshal(ud)
	if err != nil {
		logger.Error("Error to get JSON valu from user", err)
		return "", err
	}
	return string(value), nil
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) EncryptPassword() error {
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
