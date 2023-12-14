package converter

import (
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	user_entity "github.com/bellacbs/authentication-go/src/model/user/repository/entity"
)

func ConvertEntityToDomain(
	entity user_entity.UserEntity,
) user_model.UserDomainInterface {
	domain := user_model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
