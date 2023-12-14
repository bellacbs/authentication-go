package user_repository

import (
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain user_model.UserDomainInterface) (user_model.UserDomainInterface, *rest_errors.RestError)
	FindUserByID(string) (user_model.UserDomainInterface, *rest_errors.RestError)
	FindUserByEmail(string) (user_model.UserDomainInterface, *rest_errors.RestError)
}
