package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"vuegolang/auth"
	"vuegolang/helper"
	"vuegolang/user"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	// create jwt for user authenticate
	token, er := h.authService.GenerateToken(data.Id)

	if er != nil {
		errorMessages := gin.H{
			"errors": er.Error(),
		}
		return helper.ApiResponse("failed to login user", http.StatusBadRequest, "error", errorMessages)
	}

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

	if isEmailAvailable == false {
		message = false
		return helper.ApiResponse("Email is already exist", http.StatusUnprocessableEntity, "error", gin.H{"is_available": message})
	}

	return helper.ApiResponse("Email Available to use", http.StatusOK, "success", gin.H{"is_available": true})
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		response := helper.ApiResponse("failed to upload avatar", http.StatusBadRequest, "err", gin.H{"is_uploaded": false})
		c.JSON(http.StatusBadRequest, response)
	}
	
	// id from jwt
	userID := c.MustGet("currentUser").(user.User).Id

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	c.SaveUploadedFile(file, path)

	_, err = h.userService.SaveAvatar(int(userID), path)

	if err != nil {
		os.Remove(path)
		response := helper.ApiResponse("failed to upload avatar", http.StatusBadRequest, "err", gin.H{"is_uploaded": false})
		c.JSON(http.StatusBadRequest, response)
	}

	response := helper.ApiResponse("Success to upload avatar", http.StatusOK, "success", gin.H{"is_uploaded": true, "file_url": path})
	c.JSON(http.StatusOK, response)
}
