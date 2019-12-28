package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rendyfebry/go-graphql-example/controllers"
)

func main() {
	var port = flag.String("p", "8080", "Server port")
	flag.Parse()

	fmt.Println(fmt.Sprintf("Server listening on: %s", *port))

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/graphql", controllers.GraphQLHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), r)
}
