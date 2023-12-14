package user_service

import (
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	user_repository "github.com/bellacbs/authentication-go/src/model/user/repository"
)

func NewUserDomainService(
	userRepository user_repository.UserRepository,
) UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepository user_repository.UserRepository
}

type UserDomainService interface {
	CreateUser(user_model.UserDomainInterface) (user_model.UserDomainInterface, *rest_errors.RestError)
	FindUserByID(string) (user_model.UserDomainInterface, *rest_errors.RestError)
	FindUserByEmail(string) (user_model.UserDomainInterface, *rest_errors.RestError)
	UpdateUser(user_model.UserDomainInterface, string) *rest_errors.RestError
	DeleteUser(string) *rest_errors.RestError
}
