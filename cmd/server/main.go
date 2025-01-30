package main

import (
	"log"
	"net/http"

	"github.com/dfakhrizaman/ruangketiga/internal/database"
	"github.com/dfakhrizaman/ruangketiga/internal/handlers"
	"github.com/dfakhrizaman/ruangketiga/internal/places"
)

func main() {
	// 1. Connect to database
	log.Println("Starting")
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v\n", err)
	}
	defer db.Close()

	// 2. Initialize repository
	placesRepo := places.NewRepository(db)

	// 3. Initialize handler
	placesHandler := handlers.NewPlacesHandler(placesRepo)

	// 4. Wire routes
	// For simplicity, route everything under /places to PlacesHandler
	http.Handle("/places", placesHandler)
	http.Handle("/places/", placesHandler)

	// 5. Start server
	log.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}
}
