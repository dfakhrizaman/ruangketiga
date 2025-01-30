package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/dfakhrizaman/ruangketiga/internal/places"
)

type PlacesHandler struct {
	repo *places.Repository
}

func NewPlacesHandler(repo *places.Repository) *PlacesHandler {
	return &PlacesHandler{repo: repo}
}

func (h *PlacesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch r.Method {
	case http.MethodGet:
		// GET /places or GET /places/{id}
		if strings.HasPrefix(path, "/places/") {
			h.GetByID(w, r)
		} else {
			h.GetAll(w, r)
		}
	case http.MethodPost:
		// POST /places
		h.Create(w, r)
	case http.MethodPut:
		// PUT /places/{id}
		h.Update(w, r)
	case http.MethodDelete:
		// DELETE /places/{id}
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *PlacesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p places.Place
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&p); err != nil {
		log.Println("Create place error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (h *PlacesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	placesList, err := h.repo.GetAll()
	if err != nil {
		log.Println("GetAll places error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(placesList)
}

func (h *PlacesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Missing place ID", http.StatusBadRequest)
		return
	}
	id := segments[2]

	p, err := h.repo.GetByID(id)
	if err != nil {
		log.Println("GetByID error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if p == nil {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *PlacesHandler) Update(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Missing place ID", http.StatusBadRequest)
		return
	}
	id := segments[2]

	var p places.Place
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.repo.Update(id, &p); err != nil {
		log.Println("Update place error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (h *PlacesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Missing place ID", http.StatusBadRequest)
		return
	}
	id := segments[2]

	if err := h.repo.Delete(id); err != nil {
		log.Println("Delete place error:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
