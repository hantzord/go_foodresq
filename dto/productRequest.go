package dto

type ProductRequest struct {
	ProductName      string `json:"product_name"`
	Quantity         int    `json:"quantity"`
	ProductCategory  string `json:"product_category"`
	ProductPrice     int    `json:"product_price"`
	ProductCondition string `json:"product_condition"`
	ProductImage     string `json:"product_image"`
	ExpiryDate       string `json:"expiry_date"`
	Description      string `json:"description"`
}

type UpdateProductRequest struct {
	ProductName      string `json:"product_name"`
	Quantity         int    `json:"quantity"`
	ProductCategory  string `json:"product_category"`
	ProductPrice     int    `json:"product_price"`
	ProductCondition string `json:"product_condition"`
	ProductImage     string `json:"product_image"`
	ExpiryDate       string `json:"expiry_date"`
	Description      string `json:"description"`
}

type ProductDeleteRequest struct {
	ProductName string `json:"product_name"`
}
