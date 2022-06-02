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

// GetUserCart godoc
// @Summary Get User Cart
// @ID get-user-cart
// @Tags Cart
// @Produce json
// @Param Authorization header string true "Access Token"
// @Success 200 {array} carts.CartFormatter
// @Router /carts [get]
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

// GetUserCart godoc
// @Summary Save or update cart
// @ID save-or-update-cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Param request_body body carts.SaveCartInput true "Save Cart Input"
// @Success 200 {object} carts.CartFormatter
// @Router /carts [post]
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

// RemoveProduct godoc
// @Summary Remove one product from user's cart
// @ID remove-product-cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Param request_body body carts.RemoveProductInput true "Remove Product Cart Input"
// @Success 200 {object} bool
// @Router /carts/product [delete]
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

// RemoveCart godoc
// @Summary Remove the whole cart
// @ID remove-cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Success 200 {object} bool
// @Router /carts [delete]
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
