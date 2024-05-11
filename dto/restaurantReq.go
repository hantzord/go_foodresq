package dto

type RestSignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RestLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RestaurantInfoRequest struct {
	Name               string      `json:"name"`
	Latitude           float64     `json:"latitude"`
	Longitude          float64     `json:"longitude"`
	City               string      `json:"city"`
	Address            string      `json:"address"`
	PhoneNumber        string      `json:"phone_number"`
	Description        string      `json:"description"`
	RestaurantProducts interface{} `json:"restaurant_products"`
}

type RestaurantInfoUpdateRequest struct {
	Name               string      `json:"name"`
	Latitude           float64     `json:"latitude"`
	Longitude          float64     `json:"longitude"`
	City               string      `json:"city"`
	Address            string      `json:"address"`
	PhoneNumber        string      `json:"phone_number"`
	Description        string      `json:"description"`
	RestaurantProducts interface{} `json:"restaurant_products"`
}
