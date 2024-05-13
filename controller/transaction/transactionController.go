package transaction

import (
	"foodresq/dto"
	"foodresq/entities"
	errorhandler "foodresq/errorHandler"
	"foodresq/helpers"
	"foodresq/repository/transaction"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	Repo transaction.TransactionInterface
}

func NewTransactionController(repo transaction.TransactionInterface) *TransactionController {
	return &TransactionController{
		Repo: repo,
	}
}

func (tc *TransactionController) PurchaseProductByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get user ID from context or token
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		userID := uint(claims["userid"].(float64))

		// Binding request body to ProductPurchaseRequest struct
		purchaseReq := dto.PurchaseRequest{}
		if err := c.Bind(&purchaseReq); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		// Calling repository to get product by name
		product, err := tc.Repo.GetProductByName(purchaseReq.ProductName)
		if err != nil {
			return c.JSON(http.StatusNotFound, errorhandler.ResponseWriter(err.Error()))
		}

		// Create new transaction entity
		transaction := entities.Transaction{
			UserID:              userID,
			RestaurantProductID: product.ID,
			TotalAmount:         float64(product.ProductPrice) * float64(purchaseReq.Quantity),
			UniqueCode:          helpers.GenerateUniqueCode(),
			PaymentType:         "Cash",
			PickupTime:          time.Now(),
		}

		// Calling repository to insert transaction
		if err := tc.Repo.InsertTransaction(transaction); err != nil {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
		}

		// Constructing success response
		response := dto.TransactionResponse{
			Message: "Product purchased successfully",
		}

		return c.JSON(http.StatusOK, response)
	}
}
