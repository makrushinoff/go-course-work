package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lab3/mocks"
	"lab3/model"
	"testing"
)

func TestUserService_AuthenticateUser(t *testing.T) {
	t.Run("Authenticate user successfully", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := UserService{
			UserRepository: repository,
		}

		user := model.User{
			ID:       1,
			Login:    "login",
			Password: "$2a$10$zk03y2kTWecAjkdjPn5PVOVzastr5/TYmc2N/ovpJ9h.TJfBDPiOq",
		}

		userDto := model.UserDto{
			Login:    "login",
			Password: "password",
		}

		repository.On("GetByLogin", userDto.Login).Return(user, nil)

		actual := testInstance.AuthenticateUser(userDto)

		assert.NotNil(t, actual)
	})

	t.Run("Error when not found user", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := UserService{
			UserRepository: repository,
		}

		user := model.User{
			ID:       1,
			Login:    "login",
			Password: "$2a$10$zk03y2kTWecAjkdjPn5PVOVzastr5/TYmc2N/ovpJ9h.TJfBDPiOq",
		}

		userDto := model.UserDto{
			Login:    "login",
			Password: "password",
		}

		repository.On("GetByLogin", userDto.Login).Return(user, assert.AnError)

		assert.Panics(t, func() { testInstance.AuthenticateUser(userDto) })
	})

	t.Run("Error when invalid password", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := UserService{
			UserRepository: repository,
		}

		user := model.User{
			ID:       1,
			Login:    "login",
			Password: "$2a$10$zk03y2kTWecAjkdjPn5PVOVzastr5/TYmc2N/ovpJ9h.TJfBDPiOq",
		}

		userDto := model.UserDto{
			Login:    "login",
			Password: "password1",
		}

		repository.On("GetByLogin", userDto.Login).Return(user, nil)

		assert.Panics(t, func() { testInstance.AuthenticateUser(userDto) })
	})
}

func TestUserService_CreateUser(t *testing.T) {
	t.Run("Create user successfully", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := UserService{
			UserRepository: repository,
		}
		userDto := model.UserDto{
			Login:    "login",
			Password: "password",
		}

		repository.On("CreateUser", mock.Anything).Return(nil)

		testInstance.CreateUser(userDto)
	})

	t.Run("Error when create user", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := UserService{
			UserRepository: repository,
		}
		userDto := model.UserDto{
			Login:    "login",
			Password: "password",
		}

		repository.On("CreateUser", mock.Anything).Return(assert.AnError)

		assert.Panics(t, func() { testInstance.CreateUser(userDto) })
	})
}
