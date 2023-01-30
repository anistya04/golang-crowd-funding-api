package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vuegolang/helper"
	"vuegolang/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) helper.Response {

	var input user.RegisterInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ErrorResponse(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		return helper.ApiResponse("failed to register user", http.StatusUnprocessableEntity, "error", errorMessages)
	}

	data, err := h.userService.RegisterInput(input)

	if err != nil {
		return helper.ApiResponse("failed to register user", http.StatusBadRequest, "error", nil)
	}

	// to do
	// create jwt for user authenticate
	token := "wigwags"

	formatter := user.JsonFormat(data, token)

	return helper.ApiResponse("user successfully registered", 201, "success", formatter)
}

func (h *userHandler) Login(c *gin.Context) helper.Response {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ErrorResponse(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		return helper.ApiResponse("failed to login user", http.StatusUnprocessableEntity, "error", errorMessages)
	}

	data, err := h.userService.Login(input)

	if err != nil {
		return helper.ApiResponse("failed to login user", http.StatusBadRequest, "error", nil)
	}

	// to do
	// create jwt for user authenticate
	token := "wigwags"

	formatter := user.JsonFormat(data, token)

	return helper.ApiResponse("user successfully logged in", 200, "success", formatter)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) helper.Response {
	var input user.UniqueEmailInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ErrorResponse(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		return helper.ApiResponse("failed to register user", http.StatusUnprocessableEntity, "error", errorMessages)
	}

	isEmailAvailable, err := h.userService.CheckExistedUserByEmail(input)
	if err != nil {
		errors := helper.ErrorResponse(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		return helper.ApiResponse("failed to register user", http.StatusUnprocessableEntity, "error", errorMessages)
	}

	message := true
	data := gin.H{
		"is_available": message,
	}

	if isEmailAvailable == false {
		message = false
		return helper.ApiResponse("Email is already exist", http.StatusUnprocessableEntity, "error", data)
	}

	return helper.ApiResponse("Email Available to use", http.StatusOK, "success", data)
}
