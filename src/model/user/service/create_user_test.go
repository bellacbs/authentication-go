package user_service

import (
	"os"
	"testing"

	rest_errors "github.com/bellacbs/authentication-go/src/configuration/rest_erros"
	user_model "github.com/bellacbs/authentication-go/src/model/user"
	"github.com/bellacbs/authentication-go/src/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_CreateUserServices(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repository := mocks.NewMockUserRepository(controller)
	service := NewUserDomainService(repository)

	t.Run("when_user_already_exists_returns_error", func(mt *testing.T) {
		os.Setenv(user_model.COST, "10")
		defer os.Clearenv()
		id := primitive.NewObjectID().Hex()

		userDomain := user_model.NewUserDomain("test@test.com", "test", "test")
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, token, err := service.CreateUser(userDomain)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account")
	})

	t.Run("when_user_is_not_registered_returns_error", func(mt *testing.T) {
		os.Setenv(user_model.COST, "10")
		defer os.Clearenv()
		id := primitive.NewObjectID().Hex()

		userDomain := user_model.NewUserDomain("test@test.com", "test", "test")
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_errors.NewInternalServerError("Error to create user on database"))

		user, token, err := service.CreateUser(userDomain)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error to create user on database")
	})

	t.Run("when_user_is_not_registered_returns_success", func(mt *testing.T) {
		os.Setenv(user_model.COST, "10")
		os.Setenv(user_model.JWT_SECRET_KEY, "1111111122222223333333")
		defer os.Clearenv()
		id := primitive.NewObjectID().Hex()

		userDomain := user_model.NewUserDomain("test@test.com", "test", "test")
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, token, err := service.CreateUser(userDomain)

		assert.Nil(t, err)
		assert.NotNil(t, token)
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetId(), userDomain.GetId())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())

	})
}
