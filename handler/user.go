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
