package user_repository

import (
	"context"
	"os"

	"github.com/bellacbs/authentication-go/src/configuration/logger"
	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_errors.RestError {
	logger.Info("Init deleteUser repository",
		zap.String("journey", "deleteUserRepository"))
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collectionName)
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to update user",
			err,
			zap.String("journey", "DeleteUserRepository"),
			zap.String("userId", userId),
		)
		return rest_errors.NewInternalServerError(err.Error())
	}

	logger.Info(
		"DeleteUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUserRepository"))
	return nil
}
