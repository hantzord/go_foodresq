package dto

// PurchaseRequest represents the request body for purchasing a product
type PurchaseRequest struct {
	ProductName string `json:"product_name" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
	PaymentType string `json:"payment_type" validate:"required"`
	PickupTime  string `json:"pickup_time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}
