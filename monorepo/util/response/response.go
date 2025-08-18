package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns responde in json
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent && data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println("JSON encode error:", err)
		}
	}
}

// Error returns in JSON
func Erro(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"erro"`
	}{
		Err: err.Error(),
	})
}
