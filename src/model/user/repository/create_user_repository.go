package user_repository

import (
	"context"
	"os"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain user_model.UserDomainInterface) (user_model.UserDomainInterface, *rest_errors.RestError) {
	logger.Info("Init create user  repository")
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)
	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	userDomain.SetID(result.InsertedID.(string))
	return nil, nil
}
