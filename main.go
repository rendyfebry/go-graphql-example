package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rendyfebry/go-graphql-example/controllers"
)

// MakeHTTPHandler ...
func MakeHTTPHandler() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Route Declaration
	r.Post("/graphql", controllers.GraphQLHandler)

	return r
}

func main() {
	var host = flag.String("host", "localhost", "Server host")
	var port = flag.String("port", "8080", "Server port")
	flag.Parse()

	// Build the http server
	h := MakeHTTPHandler()
	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", *host, *port),
		Handler: h,

		// Set timeout
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		fmt.Println("Server Started!")
		fmt.Println(fmt.Sprintf("HTTP url: http://%s:%s", *host, *port))
		s.ListenAndServe()
	}()

	// Gracefully shutdown
	c := make(chan os.Signal, 1)
	// catch quit via SIGINT (Ctrl+C) and SIGTERM (Kill command by the OS)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	s.Shutdown(ctx)
	fmt.Println("Shutting down Mjolnir service!")
	os.Exit(0)
}
