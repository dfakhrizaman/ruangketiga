package main

import (
	"log"
	"net/http"

	"github.com/dfakhrizaman/ruangketiga/internal/database"
	"github.com/dfakhrizaman/ruangketiga/internal/handlers"
	"github.com/dfakhrizaman/ruangketiga/internal/places"
	"github.com/go-chi/chi/v5"
)

func main() {
	// 1. Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v\n", err)
	}

	defer db.Close()

	// 2. Initialize repository
	placesRepo := places.NewRepository(db)

	// 3. Initialize handler
	placesHandler := handlers.NewPlacesHandler(placesRepo)
	

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Launched"))
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pinged!"))
	})

	// Group routes under /api/v1
	r.Route("/api/v1", func(r chi.Router) {
		// /api/v1/places
		r.Post("/places", placesHandler.Create)
		r.Get("/places", placesHandler.GetAll)
		r.Get("/places/{id}", placesHandler.GetByID)
		r.Put("/places/{id}", placesHandler.Update)
		r.Delete("/places/{id}", placesHandler.Delete)
	})

	// 5. Start server
	log.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v\n", err)
	}
}
