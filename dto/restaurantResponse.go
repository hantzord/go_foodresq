package dto

type RestaurantResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type SignupRestaurantsResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RestaurantLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RestaurantProfileResponseMsg struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RestaurantInfoResponse struct {
	ID                 uint        `json:"id"`
	Name               string      `json:"name"`
	Latitude           float64     `json:"latitude"`
	Longitude          float64     `json:"longitude"`
	City               string      `json:"city"`
	Address            string      `json:"address"`
	PhoneNumber        string      `json:"phone_number"`
	Description        string      `json:"description"`
	RestaurantProducts interface{} `json:"restaurant_products"`
}

type RestaurantInfoResponseMsg struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AllRestaurantListResponse struct {
	ID                 uint        `json:"id"`
	Name               string      `json:"name"`
	Latitude           float64     `json:"latitude"`
	Longitude          float64     `json:"longitude"`
	City               string      `json:"city"`
	Address            string      `json:"address"`
	PhoneNumber        string      `json:"phone_number"`
	Description        string      `json:"description"`
	RestaurantProducts interface{} `json:"restaurant_products"`
}

type AllRestaurantListResponseMsg struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
