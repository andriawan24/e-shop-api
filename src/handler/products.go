package handler

import (
	"e-shop/src/products"
	"e-shop/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService products.Service
}

func NewProductHandler(productService products.Service) *productHandler {
	return &productHandler{productService}
}

// GetProducts godoc
// @Summary get all products
// @ID get-all-products
// @Tags Products
// @Produce json
// @Success 200 {object} products.Product
// @Router /products [get]
func (h *productHandler) GetProducts(c *gin.Context) {
	productList, err := h.productService.GetAllProducts()
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

	formatter := products.FormatProducts(productList)

	response := utils.APIResponse(
		"Success get products",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}
