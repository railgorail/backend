package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/railgorail/web-backend/internal/handlers"
	"github.com/railgorail/web-backend/internal/services"
	"github.com/railgorail/web-backend/internal/storage"
)

func main() {
	// Create instances of the layers
	storage := storage.New()
	service := services.New(storage)
	handler := handlers.New(service)

	// Create a new chi router
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define the routes
	r.Get("/", handler.HandleHello)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
