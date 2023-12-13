package user_service

import (
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
)

func (ud *userDomainService) UpdateUser(userDomain user_model.UserDomainInterface, userId string) *rest_errors.RestError {
	return nil
}
