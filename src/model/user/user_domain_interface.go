package user_model

type UserDomainInterface interface {
	GetId() string
	GetEmail() string
	GetPassword() string
	GetName() string

	SetID(string)

	EncryptPassword() error
}

func NewUserDomain(email, password, name string) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
	}
}
