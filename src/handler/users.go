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

// Register godoc
// @Summary Create an account
// @ID register-user
// @Tags Account
// @Accept json
// @Produce json
// @Param request_body body users.RegisterInput true "Register user input"
// @Success 200 {object} users.User
// @Router /sign-up [post]
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

// Login godoc
// @Summary Sign in to an account
// @ID login-user
// @Tags Account
// @Accept json
// @Produce json
// @Param request_body body users.LoginInput true "Login input"
// @Success 200 {object} users.User
// @Router /sign-in [post]
func (h *userHandler) Login(c *gin.Context) {
	// User input (email and password)
	var input users.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidationErrors(err)
		errMessage := gin.H{"errors": errors}

		response := utils.APIResponse(
			"Login failed",
			http.StatusUnprocessableEntity,
			"error",
			errMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)
	if err != nil {
		errMessage := gin.H{"errors": err.Error()}

		response := utils.APIResponse(
			"Login failed",
			http.StatusUnprocessableEntity,
			"error",
			errMessage,
		)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID)

	if err != nil {
		response := utils.APIResponse(
			"Login failed",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(loggedInUser, token)

	response := utils.APIResponse(
		"Login Success",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}

// Fetch godoc
// @Summary Get detail of an account
// @ID fetch-user
// @Tags Account
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Success 200 {object} users.User
// @Router /me [get]
func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	formatter := users.FormatUser(currentUser, "")
	response := utils.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var input users.UpdateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		fmt.Println(err.Error())
		errors := utils.FormatValidationErrors(err)
		errMsg := gin.H{"errors": errors}

		response := utils.APIResponse(
			"Failed to update user",
			http.StatusBadRequest,
			"error",
			errMsg,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	input.ID = currentUser.ID

	user, err := h.userService.UpdateUser(input)
	if err != nil {
		fmt.Println(err.Error())
		response := utils.APIResponse(
			"Failed to update user",
			http.StatusBadRequest,
			"error",
			err,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse(
		"Success registered",
		http.StatusOK,
		"success",
		user,
	)

	c.JSON(http.StatusOK, response)
}
