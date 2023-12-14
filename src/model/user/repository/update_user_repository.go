package user_repository

import (
	"context"
	"os"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"github.com/bellacbs/authentication-go/src/model/user/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUserById(userId string, userDomain user_model.UserDomainInterface) *rest_errors.RestError {
	logger.Info("Init updateUser repository",
		zap.String("journey", "updateUserRepository"))
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)
	userEntity := converter.ConvertDomainToEntity(userDomain)
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: userEntity}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user",
			err,
			zap.String("journey", "updateUserRepository"),
			zap.String("useriId", userId),
			zap.String("userName", userDomain.GetName()))
		return rest_errors.NewInternalServerError(err.Error())
	}

	logger.Info(
		"UpdateUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("userName", userEntity.Name),
		zap.String("journey", "updateUserRepository"))
	return nil
}
