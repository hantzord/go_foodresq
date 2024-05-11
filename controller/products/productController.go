package products

import (
	"foodresq/dto"
	"foodresq/entities"
	errorhandler "foodresq/errorHandler"
	"foodresq/repository/products"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	Repo products.ProductInterface
}

func NewProductController(repo products.ProductInterface) *ProductController {
	return &ProductController{
		Repo: repo,
	}
}

func (pc ProductController) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restaurantId := uint(claims["restaurantid"].(float64))

		// Binding request body to ProductRequest struct
		productReq := dto.ProductRequest{}
		if err := c.Bind(&productReq); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		// Parsing ProductCategory and ProductCondition from request
		category, err := entities.ParseProductCategory(productReq.ProductCategory)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter(err.Error()))
		}

		condition, err := entities.ParseProductCondition(productReq.ProductCondition)
		if err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter(err.Error()))
		}

		var expiryDate time.Time
		if productReq.ExpiryDate != "" {
			expiryDate, err = time.Parse("2006-01-02", productReq.ExpiryDate)
			if err != nil {
				return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid expiry date format"))
			}
		}

		// Creating Product entity
		newProduct := entities.RestaurantProduct{
			RestaurantInfoID: restaurantId,
			ProductName:      productReq.ProductName,
			ProductCategory:  category,
			ProductPrice:     productReq.ProductPrice,
			ProductCondition: condition,
			ProductImage:     productReq.ProductImage,
			ExpiryDate:       expiryDate,
			Description:      productReq.Description,
		}

		// Calling repository to create product
		if _, err := pc.Repo.CreateProduct(newProduct); err != nil {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
		}

		// Constructing success response
		response := dto.ProductResponseMsg{
			Message: "Product created successfully",
			Data:    newProduct,
		}

		return c.JSON(http.StatusCreated, response)
	}
}
