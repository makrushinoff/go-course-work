package middleware

import (
	"fmt"
	"lab3/model"
	"lab3/repository"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthenticationMiddleware struct {
	UserRepository repository.IUserRepository
}

func (am AuthenticationMiddleware) RequireAuth(c *gin.Context) {
	authorization := c.Request.Header["Authorization"]
	if authorization == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if _, err := am.authenticate(authorization[0]); err == nil {
		if apiKey := c.Request.Header["Xxx-Api-Key"]; apiKey != nil {
			if apiKey[0] == os.Getenv("API_KEY") {
				c.Next()
			} else {c.AbortWithStatus(http.StatusUnauthorized)}
		} else {c.AbortWithStatus(http.StatusUnauthorized)}
	} else {c.AbortWithStatus(http.StatusUnauthorized)}
}

func (am AuthenticationMiddleware) authenticate(tokenString string) (model.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		secret := os.Getenv("SECRET")
		bytes := []byte(secret)
		return bytes, nil
	})
	if err != nil {
		return model.User{}, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	currentTime := time.Now()
	timeLong := currentTime.Unix()
	expirationTime := claims["exp"]
	currentTimeFloat := float64(timeLong)
	expirationTimeFloat := expirationTime.(float64)
	if currentTimeFloat > expirationTimeFloat {
		return model.User{}, fmt.Errorf("token is expired")
	}
	login := claims["sub"]
	loginString := login.(string)
	userEntity, err := am.UserRepository.GetByLogin(loginString)
	if err != nil {
		return model.User{}, err
	}
	return userEntity, nil
}
