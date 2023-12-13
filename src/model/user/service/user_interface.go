package user_service

import (
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(user_model.UserDomainInterface) *rest_errors.RestError
	UpdateUser(user_model.UserDomainInterface, string) *rest_errors.RestError
	FindUser(string) (*user_model.UserDomainInterface, *rest_errors.RestError)
	DeleteUser(string) *rest_errors.RestError
}
