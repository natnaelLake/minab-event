package main

import (
	"fmt"
	"log"
	"net/http"
	// "os"

	"backend/controller"
	"backend/middlewares"
)

func main() {

	// Define HTTP server and routes
	http.HandleFunc("/register", middlewares.ApplyMiddleware(controller.Signup, middlewares.Logger, middlewares.CorsMiddleware))
	http.HandleFunc("/login", middlewares.ApplyMiddleware(controller.Login, middlewares.Logger, middlewares.CorsMiddleware))
	http.HandleFunc("/uploadImage", middlewares.ApplyMiddleware(controller.UploadImagesHandler, middlewares.Logger, middlewares.CorsMiddleware))
	http.HandleFunc("/process-payment", middlewares.ApplyMiddleware(controller.ProcessPaymentHandler, middlewares.Logger, middlewares.CorsMiddleware))
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	})
	// Start HTTP server
	fmt.Println("Server is running on port 5050")
	log.Fatal(http.ListenAndServe(":5050", nil))
}
