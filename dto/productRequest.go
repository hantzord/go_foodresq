package dto

type ProductRequest struct {
	ProductName      string `json:"product_name"`
	ProductCategory  string `json:"product_category"`
	ProductPrice     int    `json:"product_price"`
	ProductCondition string `json:"product_condition"`
	ProductImage     string `json:"product_image"`
	ExpiryDate       string `json:"expiry_date"`
	Description      string `json:"description"`
}
