package dto

type TransactionResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
