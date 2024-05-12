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
		restaurantID := uint(claims["restaurantid"].(float64))

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

		// Parse ExpiryDate to time.Time
		var expiryDate time.Time
		if productReq.ExpiryDate != "" {
			// Parse tanggal kedaluwarsa dari string ke time.Time
			expiryDate, err = time.Parse("2006-01-02", productReq.ExpiryDate)
			if err != nil {
				return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid expiry date format"))
			}
			// Konversi ke Waktu Indonesia Barat (WIB)
			loc, err := time.LoadLocation("Asia/Jakarta")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
			}
			expiryDate = expiryDate.In(loc)
		}

		// Creating Product entity
		newProduct := entities.RestaurantProduct{
			RestaurantInfoID: restaurantID,
			ProductName:      productReq.ProductName,
			Quantity:         productReq.Quantity,
			ProductCategory:  category,
			ProductPrice:     productReq.ProductPrice,
			ProductCondition: condition,
			ProductImage:     productReq.ProductImage,
			ExpiryDate:       expiryDate,
			Description:      productReq.Description,
		}

		// Calling repository to create product
		createdProduct, err := pc.Repo.CreateProduct(restaurantID, newProduct)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
		}

		// Constructing success response
		response := dto.ProductResponseMsg{
			Message: "Product created successfully",
			Data:    createdProduct,
		}

		return c.JSON(http.StatusCreated, response)
	}
}

func (pc ProductController) GetAllProduct() echo.HandlerFunc {

	return func(c echo.Context) error {
		products, err := pc.Repo.GetAllProduct()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
		}

		response := dto.ProductListResponseMsg{
			Message: "List of all products",
			Data:    products,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (pc ProductController) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restaurantID := uint(claims["restaurantid"].(float64))

		// Binding request body to ProductRequest struct
		productReq := dto.UpdateProductRequest{}
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

		// Parse ExpiryDate to time.Time
		var expiryDate time.Time
		if productReq.ExpiryDate != "" {
			// Parse tanggal kedaluwarsa dari string ke time.Time
			expiryDate, err = time.Parse("2006-01-02", productReq.ExpiryDate)
			if err != nil {
				return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid expiry date format"))
			}
			// Konversi ke Waktu Indonesia Barat (WIB)
			loc, err := time.LoadLocation("Asia/Jakarta")
			if err != nil {
				return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
			}
			expiryDate = expiryDate.In(loc)
		}

		// Creating Product entity
		updatedProduct := entities.RestaurantProduct{
			RestaurantInfoID: restaurantID,
			ProductName:      productReq.ProductName,
			Quantity:         productReq.Quantity,
			ProductCategory:  category,
			ProductPrice:     productReq.ProductPrice,
			ProductCondition: condition,
			ProductImage:     productReq.ProductImage,
			ExpiryDate:       expiryDate,
			Description:      productReq.Description,
		}

		// Calling repository to update product
		updatedProduct, err = pc.Repo.UpdateProduct(restaurantID, updatedProduct)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
		}

		// Constructing success response
		response := dto.ProductResponseMsg{
			Message: "Product updated successfully",
			Data:    updatedProduct,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (pc *ProductController) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		restaurantID := uint(claims["restaurantid"].(float64))

		// Binding request body to ProductDeleteRequest struct
		deleteReq := dto.ProductDeleteRequest{}
		if err := c.Bind(&deleteReq); err != nil {
			return c.JSON(http.StatusBadRequest, errorhandler.ResponseWriter("Invalid request body"))
		}

		// Calling repository to delete product
		err := pc.Repo.DeleteProduct(restaurantID, uint(deleteReq.ProductID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, errorhandler.ResponseWriter(err.Error()))
		}

		// Constructing success response
		response := dto.ProductDeleteResponseMsg{
			Message: "Product deleted successfully",
		}

		return c.JSON(http.StatusOK, response)
	}
}
