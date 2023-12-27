package user_repository

import (
	"os"
	"testing"

	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	database_name := "user_database_test"
	collection_name := "user_collection_test"

	os.Setenv(MONGODB_USER_COLLECTION, collection_name)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_user_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := user_model.NewUserDomain("test@test.com", "test", "test")

		domain.SetID(primitive.NewObjectID().Hex())

		err := repo.UpdateUserById(domain.GetId(), domain)
		assert.Nil(t, err)
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := user_model.NewUserDomain("test@test.com", "test", "test")

		domain.SetID(primitive.NewObjectID().Hex())
		err := repo.UpdateUserById(domain.GetId(), domain)
		assert.NotNil(t, err)
	})
}
