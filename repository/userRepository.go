package repository

import (
	"lab3/database"
	"lab3/model"
)

type IUserRepository interface {
	GetByLogin(login string) (model.User, error)
	CreateUser(userDto model.UserDto) error
}

type UserRepository struct {
	DB database.IDB
}

func (ur UserRepository) GetByLogin(login string) (model.User, error) {
	entity, err := ur.DB.FindUserByLogin(login)
	if err != nil {
		return model.User{}, err
	}
	return entity, nil
}

func (ur UserRepository) CreateUser(userDto model.UserDto) error {
	user := model.User{Login: userDto.Login, Password: userDto.Password}
	createErr := ur.DB.CreateUser(user)
	if createErr != nil {
		return createErr
	}

	return nil
}
