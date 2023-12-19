package user_repository

import (
	"fmt"
	"os"
	"testing"

	user_entity "github.com/bellacbs/authentication-go/src/model/user/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv(MONGODB_USER_COLLECTION, collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()
	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("Return_success_when_email_sending", func(mt *mtest.T) {
		userEntity := user_entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity)))
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetId(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
	})

	mtestDb.Run("return_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch))
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserById(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv(MONGODB_USER_COLLECTION, collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("return_success_sending_a_valid_id", func(mt *mtest.T) {
		userEntity := user_entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "test",
			Name:     "test",
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			convertEntityToBson(userEntity),
		))
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID(userEntity.ID.Hex())
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetId(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
	})

	mtestDb.Run("return_error_when_mongodb_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID("test")
		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDb.Run("return_no_document_found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
		))
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID("test")
		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "User not found with this ID test")
		assert.Nil(t, userDomain)
	})
}

func convertEntityToBson(userEntity user_entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
	}
}
