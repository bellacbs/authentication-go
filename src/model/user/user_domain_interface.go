package user_model

import rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"

type UserDomainInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetName() string

	SetID(string)
	SetName(string)

	EncryptPassword() error
	CheckPasswordHash(string, string) *rest_errors.RestError
	GenerateToken() (string, *rest_errors.RestError)
}

func NewUserDomain(email, password, name string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
	}
}

func NewUserUpdateDomain(name string) UserDomainInterface {
	return &userDomain{
		name: name,
	}
}

func NewUserLoginDomain(email, password string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}
