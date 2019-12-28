package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rendyfebry/go-graphql-example/internal/utils"
	gqlRepo "github.com/rendyfebry/go-graphql-example/repository/graphql"
	"github.com/rendyfebry/go-graphql-example/repository/models"

	"github.com/graphql-go/graphql"
)

// GraphQLHandler ...
func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res := &models.Response{
			Error: "Invalid body",
		}

		utils.SendJSONResponse(w, http.StatusInternalServerError, res)
		return
	}

	defer func() {
		io.Copy(ioutil.Discard, r.Body)
		r.Body.Close()
	}()

	var apolloQuery map[string]interface{}
	if err := json.Unmarshal(buf, &apolloQuery); err != nil {
		res := &models.Response{
			Error: "Error on Unmarshalling!!!",
		}

		utils.SendJSONResponse(w, http.StatusUnprocessableEntity, res)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        gqlRepo.UserSchema,
		RequestString: apolloQuery["query"].(string),
	})

	if len(result.Errors) > 0 {
		res := &models.Response{
			Error: fmt.Sprintf("Wrong result, unexpected errors: %v", result.Errors),
		}

		utils.SendJSONResponse(w, http.StatusBadRequest, res)
		return
	}

	res := &models.Response{
		Data: result.Data,
	}

	utils.SendJSONResponse(w, http.StatusOK, res)
	return
}
