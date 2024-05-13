package dto

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type GetResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
