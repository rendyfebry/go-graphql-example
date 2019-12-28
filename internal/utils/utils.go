package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rendyfebry/go-graphql-example/repository/models"
)

// SendJSONResponse this function will give Content-Type JSON on given response
func SendJSONResponse(w http.ResponseWriter, httpCode int, res *models.Response) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(httpCode)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
