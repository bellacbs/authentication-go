package user_repository

import (
	"context"
	"os"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"github.com/bellacbs/authentication-go/src/model/user/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(userDomain user_model.UserDomainInterface) (user_model.UserDomainInterface, *rest_errors.RestError) {
	logger.Info("Init createUser repository",
		zap.String("journey", "createUser"))
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)
	value := converter.ConvertDomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"))
	return converter.ConvertEntityToDomain(*value), nil
}
