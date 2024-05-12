package dto

type ProductResponseMsg struct {
	Message string `json:"message"`
	Data    interface{}
}

type ProductListResponseMsg struct {
	Message string `json:"message"`
	Data    interface{}
}

type ProductDeleteResponseMsg struct {
	Message string `json:"message"`
	Data    interface{}
}
