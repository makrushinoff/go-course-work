package controller

import (
	"lab3/model"
	"lab3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	UserService service.IUserService
}

// MakeCalculations	godoc
// @Summary 		register as new user
// @Description 	registers new user with provided date
// @Param 			registrationForm body model.UserDto true "Submit"
// @Produce			application/json
// @Tags 			users
// @Router			/register [post]
func (ac AuthenticationController) RegisterUser(context *gin.Context) {
	var userDto model.UserDto
	if err := context.BindJSON(&userDto); err == nil {
		ac.UserService.CreateUser(userDto)
		context.Status(http.StatusOK)
	}
}

// MakeCalculations	godoc
// @Summary 		get token
// @Description 	authenticate user in the system
// @Param 			authenticationForm body model.UserDto true "Submit"
// @Produce			application/json
// @Tags 			users
// @Success 		200 {object} model.Token
// @Router			/authenticate [post]
func (ac AuthenticationController) AuthenticateUser(context *gin.Context) {
	var userDto model.UserDto
	if err := context.BindJSON(&userDto); err == nil {
		context.IndentedJSON(http.StatusOK, ac.UserService.AuthenticateUser(userDto))
	}
}
