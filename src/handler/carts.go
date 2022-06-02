package handler

import (
	"e-shop/src/carts"
	"e-shop/src/users"
	"e-shop/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type cartsHandler struct {
	cartsService carts.Service
}

func NewCartsHandler(cartsService carts.Service) *cartsHandler {
	return &cartsHandler{cartsService}
}

func (h *cartsHandler) GetUserCart(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	userID := currentUser.ID

	cart, err := h.cartsService.GetUserCart(userID)
	if err != nil {
		fmt.Println(err.Error())
		response := utils.APIResponse(
			"Failed to get cart",
			http.StatusBadRequest,
			"error",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := carts.FormatCart(cart)

	response := utils.APIResponse(
		"Success get products",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}

func (h *cartsHandler) SaveOrUpdateCart(c *gin.Context) {
	var input carts.SaveCartInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println(err.Error())
		errors := utils.FormatValidationErrors(err)
		errMsg := gin.H{"errors": errors}

		response := utils.APIResponse(
			"Failed to update carts",
			http.StatusBadRequest,
			"error",
			errMsg,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	userId := currentUser.ID

	cart, err := h.cartsService.SaveCart(input, userId)
	if err != nil {
		response := utils.APIResponse(
			"Failed to update carts",
			http.StatusBadRequest,
			"error",
			err,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := carts.FormatCart(cart)

	response := utils.APIResponse(
		"Success update carts",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}

func (h *cartsHandler) RemoveProduct(c *gin.Context) {
	var input carts.RemoveProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println(err.Error())
		errors := utils.FormatValidationErrors(err)
		errMsg := gin.H{"errors": errors}

		response := utils.APIResponse(
			"Failed to remove product",
			http.StatusBadRequest,
			"error",
			errMsg,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	userId := currentUser.ID

	success, err := h.cartsService.RemoveProduct(userId, input)
	if err != nil {
		response := utils.APIResponse(
			"Failed to remove product",
			http.StatusBadRequest,
			"error",
			err,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse(
		"Success remove product",
		http.StatusOK,
		"success",
		success,
	)

	c.JSON(http.StatusOK, response)
}

func (h *cartsHandler) RemoveCart(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(users.User)
	userId := currentUser.ID

	success, err := h.cartsService.RemoveCart(userId)
	if err != nil {
		response := utils.APIResponse(
			"Failed to remove cart",
			http.StatusBadRequest,
			"error",
			err,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse(
		"Success remove cart",
		http.StatusOK,
		"success",
		success,
	)

	c.JSON(http.StatusOK, response)
}
