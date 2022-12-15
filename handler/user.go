package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vuegolang/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) (user.User, error) {

	var input user.RegisterInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	data, err := h.userService.RegisterInput(input)
	if err != nil {
		return data, err
	}

	return data, nil

}
