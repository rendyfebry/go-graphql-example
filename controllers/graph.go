package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rendyfebry/go-graphql-example/internal/utils"
	gqlRepo "github.com/rendyfebry/go-graphql-example/repository/graphql"

	"github.com/graphql-go/graphql"
)

// GraphQLHandler ...
func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer func() {
		io.Copy(ioutil.Discard, r.Body)
		r.Body.Close()
	}()

	var apolloQuery map[string]interface{}
	if err := json.Unmarshal(buf, &apolloQuery); err != nil {
		utils.SendJSONResponse(w, http.StatusUnprocessableEntity, "Error on Unmarshalling!!!", nil)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        gqlRepo.UserSchema,
		RequestString: apolloQuery["query"].(string),
	})

	if len(result.Errors) > 0 {
		utils.SendJSONResponse(w, http.StatusBadRequest, fmt.Sprintf("wrong result, unexpected errors: %v", result.Errors), nil)
		return
	}

	utils.SendJSONResponse(w, 0, "Success", result.Data)
	return
}
