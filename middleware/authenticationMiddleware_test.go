package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"lab3/mocks"
	"lab3/model"
	"os"
	"testing"
	"time"
)

func TestAuthenticationMiddleware_authenticate(t *testing.T) {
	t.Run("Authenticate user successfully", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := AuthenticationMiddleware{
			UserRepository: repository,
		}

		token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "login",
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		}).SignedString([]byte(os.Getenv("SECRET")))

		user := model.User{
			ID:       1,
			Login:    "login",
			Password: "password",
		}
		repository.On("GetByLogin", "login").Return(user, nil)

		actual, actualErr := testInstance.authenticate(token)

		assert.Nil(t, actualErr)
		assert.Equal(t, user, actual)
	})

	t.Run("Error when token has invalid method", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := AuthenticationMiddleware{
			UserRepository: repository,
		}

		token, _ := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
			"sub": "login",
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		}).SignedString([]byte(os.Getenv("SECRET")))

		actual, actualErr := testInstance.authenticate(token)

		assert.Error(t, actualErr)
		assert.Equal(t, model.User{}, actual)
	})

	t.Run("Error when token is not parseable", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := AuthenticationMiddleware{
			UserRepository: repository,
		}

		token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "login",
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		}).SignedString([]byte(os.Getenv("SECRET")))

		actual, actualErr := testInstance.authenticate(token + "123")

		assert.Error(t, actualErr)
		assert.Equal(t, model.User{}, actual)
	})

	t.Run("Error when token is expired", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := AuthenticationMiddleware{
			UserRepository: repository,
		}

		token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "login",
			"exp": time.Now().AddDate(-1, 0, 0).Unix(),
		}).SignedString([]byte(os.Getenv("SECRET")))

		actual, actualErr := testInstance.authenticate(token)
		assert.Error(t, actualErr)
		assert.Equal(t, model.User{}, actual)
	})

	t.Run("Error when db returns error", func(t *testing.T) {
		repository := mocks.NewIUserRepository(t)
		testInstance := AuthenticationMiddleware{
			UserRepository: repository,
		}

		token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "login",
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		}).SignedString([]byte(os.Getenv("SECRET")))

		repository.On("GetByLogin", "login").Return(model.User{}, assert.AnError)

		actual, actualErr := testInstance.authenticate(token)

		assert.Error(t, actualErr)
		assert.Equal(t, model.User{}, actual)
	})
}
