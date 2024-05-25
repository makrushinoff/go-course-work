package repository

import (
	"github.com/stretchr/testify/assert"
	"lab3/mocks"
	"lab3/model"
	"testing"
)

func TestUserRepository_CreateUser(t *testing.T) {
	t.Run("Create user successfully", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := UserRepository{DB: db}

		userDto := model.UserDto{Login: "login", Password: "password"}
		user := model.User{Login: userDto.Login, Password: userDto.Password}
		db.On("CreateUser", user).Return(nil)

		actualErr := testInstance.CreateUser(userDto)

		assert.Nil(t, actualErr)
	})

	t.Run("Error when Create user", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := UserRepository{DB: db}

		userDto := model.UserDto{Login: "login", Password: "password"}
		user := model.User{Login: userDto.Login, Password: userDto.Password}

		db.On("CreateUser", user).Return(assert.AnError)

		actualErr := testInstance.CreateUser(userDto)

		assert.Error(t, actualErr)
	})
}

func TestUserRepository_GetByLogin(t *testing.T) {
	t.Run("Get user by id successfully", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := UserRepository{DB: db}

		user := model.User{ID: 1, Login: "login", Password: "password"}

		db.On("FindUserByLogin", "login").Return(user, nil)

		actual, actualErr := testInstance.GetByLogin("login")

		assert.Nil(t, actualErr)
		assert.Equal(t, user, actual)
	})

	t.Run("Error when get user by id", func(t *testing.T) {
		db := mocks.NewIDB(t)
		testInstance := UserRepository{DB: db}

		user := model.User{ID: 1, Login: "login", Password: "password"}

		db.On("FindUserByLogin", "login").Return(user, assert.AnError)

		actual, actualErr := testInstance.GetByLogin("login")

		assert.Error(t, actualErr)
		assert.Equal(t, model.User{}, actual)
	})
}
