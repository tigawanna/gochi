package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(basicHandler),
	}

	fmt.Println("Starting server...")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
