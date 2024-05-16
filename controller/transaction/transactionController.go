package transaction

import (
	"foodresq/dto"
	"foodresq/entities"
	errorhandler "foodresq/errorHandler"
	"foodresq/helpers"
	"foodresq/repository/transaction"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	Repo      transaction.TransactionInterface
	Validator *validator.Validate
}

func NewTransactionController(repo transaction.TransactionInterface) *TransactionController {
	return &TransactionController{
		Repo:      repo,
		Validator: validator.New(),
	}
}

// func (tc *TransactionController) PurchaseProductByName() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// Get user ID from context or token
// 		uid := c.Get("user").(*jwt.Token)
// 		claims := uid.Claims.(jwt.MapClaims)
// 		userID := uint(claims["userid"].(float64))

// 		// Binding request body to ProductPurchaseRequest struct
// 		purchaseReq := dto.PurchaseRequest{}
// 		if err := c.Bind(&purchaseReq); err != nil {
// 			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
// 		}

// 		// Calling repository to get product by name
// 		product, err := tc.Repo.GetProductByName(purchaseReq.ProductName)
// 		if err != nil {
// 			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))
// 		}

// 		// Create new transaction entity
// 		transaction := entities.Transaction{
// 			UserID:              userID,
// 			RestaurantProductID: product.ID,
// 			TotalAmount:         float64(product.ProductPrice) * float64(purchaseReq.Quantity),
// 			UniqueCode:          helpers.GenerateUniqueCode(),
// 			PaymentType:         "Cash",
// 			PickupTime:          time.Now(),
// 		}

// 		// Calling repository to insert transaction
// 		if err := tc.Repo.InsertTransaction(transaction); err != nil {
// 			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
// 		}

// 		// Constructing success response
// 		response := dto.TransactionResponse{
// 			Message: "Product purchased successfully",
// 		}

// 		return c.JSON(http.StatusOK, response)
// 	}
// }

func (tc *TransactionController) PurchaseProductByName() echo.HandlerFunc {
	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := uint(claims["userid"].(float64))

		// Bind request body to DTO
		var transactionReq dto.PurchaseRequest
		if err := c.Bind(&transactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		// Validate request data
		if err := tc.Validator.Struct(&transactionReq); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter(err.Error()))
		}

		// Get product by name
		product, err := tc.Repo.GetProductByName(transactionReq.ProductName)
		if err != nil {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))
		}

		// Calculate total amount
		totalAmount := int(transactionReq.Quantity) * product.ProductPrice
		totalAmountFloat := float64(totalAmount)

		// Parse pickup time
		pickupTime, err := time.Parse(time.RFC3339, transactionReq.PickupTime)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid pickup time format"))
		}

		// Generate unique code
		uniqueCode := helpers.GenerateUniqueCode()

		// Create transaction entity
		newTransaction := entities.NewTransaction(
			userID,
			product.ID,
			totalAmountFloat,
			uniqueCode,
			transactionReq.PaymentType,
			pickupTime,
		)

		// Save transaction
		createdTransaction, err := tc.Repo.CreateTransaction(*newTransaction)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter("Failed to create transaction"))
		}

		// Return response
		return c.JSON(http.StatusCreated, createdTransaction)
	}
}

func (tc *TransactionController) ValidateUniqueCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract restaurant ID from JWT
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restaurantID := uint(claims["restaurantid"].(float64))

		uniqueCode := c.Param("uniqueCode")

		// Get transaction by unique code
		transaction, err := tc.Repo.GetTransactionByUniqueCode(uniqueCode)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{"status": false, "message": "Transaction not found"})
		}

		// Get product associated with the transaction to verify restaurant ownership
		product, err := tc.Repo.GetProductByID(transaction.RestaurantProductID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "message": "Failed to get product details"})
		}

		// Verify that the product belongs to the restaurant
		if product.RestaurantInfoID != restaurantID {
			return c.JSON(http.StatusForbidden, map[string]interface{}{"status": false, "message": "You do not have permission to view this transaction"})
		}

		// Return transaction details
		return c.JSON(http.StatusOK, transaction)
	}
}
