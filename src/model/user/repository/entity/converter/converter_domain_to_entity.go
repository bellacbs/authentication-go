package converter

import (
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	user_entity "github.com/bellacbs/authentication-go/src/model/user/repository/entity"
)

func ConvertDomainToEntity(
	domain user_model.UserDomainInterface,
) *user_entity.UserEntity {
	return &user_entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
	}
}
