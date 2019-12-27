package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	gqlRepo "github.com/rendyfebry/go-graphql-example/repository/graphql"

	"github.com/graphql-go/graphql"
)

// GraphQLHandler ...
func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		io.Copy(ioutil.Discard, r.Body)
		r.Body.Close()
	}()

	var apolloQuery map[string]interface{}
	if err := json.Unmarshal(buf, &apolloQuery); err != nil {
		fmt.Println(err)
		fmt.Println("Error on Unmarshalling!!!")
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        gqlRepo.UserSchema,
		RequestString: apolloQuery["query"].(string),
	})

	if len(result.Errors) > 0 {
		fmt.Println(fmt.Sprintf("wrong result, unexpected errors: %v", result.Errors))
	}

	resultByte, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultByte)
}

func main() {
	http.HandleFunc("/graphql", GraphQLHandler)
	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
