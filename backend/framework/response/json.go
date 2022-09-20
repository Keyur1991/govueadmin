package response

import (
	"net/http"
	"encoding/json"
)


func Json(w *http.ResponseWriter, status int, data interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(data)
}