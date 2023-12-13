package view

import (
	"github.com/bellacbs/authentication-go/src/controller/model/response"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"github.com/google/uuid"
)

func ConvertDomainToResponse(
	userDomain user_model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    uuid.New().String(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
	}
}
