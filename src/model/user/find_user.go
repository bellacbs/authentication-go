package user_model

import rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"

func (ud *UserDomain) FindUser(userId string) (*UserDomain, *rest_errors.RestError) {
	return ud, nil
}
