package errorhandler

// func BadRequest(w http.ResponseWriter, message string) {
// 	w.WriteHeader(http.StatusBadRequest)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"error": message,
// 	})
// }

// func InternalServerError(w http.ResponseWriter, message string) {
// 	w.WriteHeader(http.StatusInternalServerError)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"error": message,
// 	})
// }

// func NotFound(w http.ResponseWriter, message string) {
// 	w.WriteHeader(http.StatusNotFound)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"error": message,
// 	})
// }

// func Unauthorized(w http.ResponseWriter, message string) {
// 	w.WriteHeader(http.StatusUnauthorized)
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"error": message,
// 	})
// }

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func ResponseWriter(message string) *ErrorResponse {
	return &ErrorResponse{
		Status:  true,
		Message: message,
	}
}
