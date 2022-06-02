package handler

import (
	"e-shop/src/carts"
	"e-shop/src/transactions"
	"e-shop/src/users"
	"e-shop/src/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transactions.Service
	cartService        carts.Service
}

func NewTransactionHandler(transactionService transactions.Service, cartService carts.Service) *transactionHandler {
	return &transactionHandler{transactionService, cartService}
}

func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)
	userId := currentUser.ID

	transactionList, err := h.transactionService.GetTransactions(userId)
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

	formatter := transactions.FormatTransactions(transactionList)

	response := utils.APIResponse(
		"Success get transactions",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CheckoutTransaction(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(users.User)
	userId := currentUser.ID

	// Ambil data cart
	cart, err := h.cartService.GetUserCart(userId)
	if err != nil {
		response := utils.APIResponse(
			"Failed to create transactions",
			http.StatusBadRequest,
			"error",
			err,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create transaction object
	var transaction transactions.Transaction

	// Get total price and transaction detail
	var transactionDetail []transactions.TransactionDetail
	var totalPrice int
	for _, product := range cart.CartDetail {
		if product.Product.DiscountedPrice > 0 {
			totalPrice += (product.Product.DiscountedPrice * product.Quantity)
		} else {
			totalPrice += (product.Product.Price * product.Quantity)
		}

		transactionDetail = append(transactionDetail, transactions.TransactionDetail{
			ProductID: product.Product.ID,
			Quantity:  product.Quantity,
		})
	}

	// Set deadline to one day
	deadline := time.Now().Add(time.Hour * 24)

	transaction.UserID = userId
	transaction.TotalPrice = totalPrice
	transaction.Deadline = deadline
	transaction.Status = "pending"

	newTransaction, err := h.transactionService.CheckoutCart(transaction, transactionDetail, currentUser)
	if err != nil {
		response := utils.APIResponse(
			"Failed to create transactions",
			http.StatusBadRequest,
			"error",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Remove cart after transactions
	_, err = h.cartService.RemoveCart(userId)
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

	formatter := transactions.FormatTransaction(newTransaction)

	response := utils.APIResponse(
		"Success checkout transaction",
		http.StatusOK,
		"success",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transactions.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := utils.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	err = h.transactionService.ProccessPayment(input)
	if err != nil {
		response := utils.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)

		return
	}

	c.JSON(http.StatusOK, input)
}
