package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", rootHandler)
	router.Get("/hello", helloHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Starting server...")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Homeroute"))
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello route"))
}
