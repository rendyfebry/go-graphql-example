package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rendyfebry/go-graphql-example/repository/models"
)

// SendJSONResponse this function will give Content-Type JSON on given response
func SendJSONResponse(w http.ResponseWriter, error interface{}, message interface{}, data interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := &models.Response{
		Error:   error,
		Data:    data,
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
