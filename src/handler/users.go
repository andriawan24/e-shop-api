package handler

import (
	"e-shop/src/auth"
	"e-shop/src/users"
	"e-shop/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService users.Service
	authService auth.Service
}

func NewUserHandler(userService users.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	// Grab user input to struct
	var input users.RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println(err.Error())
		errors := utils.FormatValidationErrors(err)
		errMsg := gin.H{"errors": errors}

		response := utils.APIResponse(
			"Failed to register",
			http.StatusBadRequest,
			"error",
			errMsg,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.userService.Register(input)
	if err != nil {
		response := utils.APIResponse(
			"Failed to register",
			http.StatusBadRequest,
			"error",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Generate token here
	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		response := utils.APIResponse(
			"Failed to register",
			http.StatusBadRequest,
			"error",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(user, token)

	response := utils.APIResponse(
		"Success registered",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}
