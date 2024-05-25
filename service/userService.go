package service

import (
	"lab3/model"
	"lab3/repository"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(userDto model.UserDto)
	AuthenticateUser(userDto model.UserDto) model.Token
}

type UserService struct {
	UserRepository repository.IUserRepository
}

func (us UserService) CreateUser(userDto model.UserDto) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), 10)
	userDto.Password = string(hashedPassword)
	err = us.UserRepository.CreateUser(userDto)
	if err != nil {
		log.Panic(err)
	}
}

func (us UserService) AuthenticateUser(userDto model.UserDto) model.Token {
	user, err := us.UserRepository.GetByLogin(userDto.Login)
	if err != nil {
		log.Panic(err)
		return model.Token{}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password))
	if err != nil {
		log.Panic(err)
		return model.Token{}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Login,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
	return model.Token{Token: tokenString}
}
