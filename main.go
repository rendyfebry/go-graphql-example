package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rendyfebry/go-graphql-example/controllers"
)

func main() {
	var port = flag.String("p", "8080", "Server port")
	flag.Parse()

	fmt.Println(fmt.Sprintf("Server listening on: %s", *port))

	r := mux.NewRouter().StrictSlash(true)
	r.Methods("POST").
		Path("/graphql").
		HandlerFunc(controllers.GraphQLHandler)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), loggedRouter)
}
