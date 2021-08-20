package handler

import (
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	// handler dependenciesnya kesiapa
	// kalau di ruby on rails handler itu controller nya
	// tidak jauh2 ke service
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// get input from user
	// map input from user to struct RegisterUserInput
	// struct above we pass as a parameter service

	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, user)
}
