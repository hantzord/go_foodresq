package dto

type PurchaseRequest struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
