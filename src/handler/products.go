package handler

import (
	"e-shop/src/products"
	"e-shop/src/utils"
	"fmt"
	"net/http"
	"strconv"

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
// @Param category_id query int false "Category ID"
// @Success 200 {array} products.Product
// @Router /products [get]
func (h *productHandler) GetProducts(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Query("category_id"))

	productList, err := h.productService.GetAllProducts(categoryId)
	if err != nil {
		fmt.Println(err.Error())
		errors := utils.FormatValidationErrors(err)
		errMsg := gin.H{"errors": errors}

		response := utils.APIResponse(
			"Failed to get products",
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

// Getcategories godoc
// @Summary get all categories
// @ID get-all-categories
// @Tags Products
// @Produce json
// @Success 200 {array} products.Category
// @Router /categories [get]
func (h *productHandler) GetCategories(c *gin.Context) {
	categoryList, err := h.productService.GetCategories()
	if err != nil {
		fmt.Println(err.Error())
		errors := utils.FormatValidationErrors(err)
		errMsg := gin.H{"errors": errors}

		response := utils.APIResponse(
			"Failed to get categories",
			http.StatusBadRequest,
			"error",
			errMsg,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := products.FormatCategories(categoryList)

	response := utils.APIResponse(
		"Success get categories",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}
