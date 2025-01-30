package places

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const testDB = "postgres://myuser:mypassword@localhost:5432/ruangketiga_test?sslmode=disable"

var testRepo *Repository

// Setup function runs before all tests.
func TestMain(m *testing.M) {
    db, err := sql.Open("postgres", testDB)
    if err != nil {
        log.Fatalf("Failed to connect to test database: %v", err)
    }
    defer db.Close()

    // Clear test data before each run
    db.Exec("DELETE FROM places")

    testRepo = NewRepository(db)

    os.Exit(m.Run()) // Run tests
}

// Test `Create()`
func TestCreatePlace(t *testing.T) {
    place := &Place{
        Name:        "Test Park",
        Type:        "park",
        Address:     "123 Test St",
        District:    "Central",
        City:        "TestCity",
        Latitude:    40.7128,
        Longitude:   -74.0060,
        SizeM2:      5000,
    }

    err := testRepo.Create(place)
    if err != nil {
        t.Fatalf("Failed to create place: %v", err)
    }

    if place.ID == "" {
        t.Errorf("Expected ID to be set, got empty string")
    }
}

// Test `GetAll()`
func TestGetAllPlaces(t *testing.T) {
    places, err := testRepo.GetAll()
    if err != nil {
        t.Fatalf("Failed to get places: %v", err)
    }
    if len(places) == 0 {
        t.Errorf("Expected at least one place, got 0")
    }
}

// Test `GetByID()`
func TestGetByID(t *testing.T) {
    place := &Place{
        Name:     "Find Me",
        Type:     "cafe",
        Address:  "Hidden Address",
        District: "West",
        City:     "HiddenCity",
    }
    testRepo.Create(place)

    found, err := testRepo.GetByID(place.ID)
    if err != nil {
        t.Fatalf("Error finding place by ID: %v", err)
    }
    if found == nil {
        t.Errorf("Expected to find a place, got nil")
    }
}

// Test `Update()`
func TestUpdatePlace(t *testing.T) {
    place := &Place{
        Name:     "Before Update",
        Type:     "restaurant",
        Address:  "Old Address",
        District: "North",
        City:     "OldCity",
    }
    testRepo.Create(place)

    place.Name = "After Update"
    place.Address = "New Address"
    err := testRepo.Update(place.ID, place)
    if err != nil {
        t.Fatalf("Error updating place: %v", err)
    }

    updated, _ := testRepo.GetByID(place.ID)
    if updated.Name != "After Update" {
        t.Errorf("Update failed: expected 'After Update', got '%s'", updated.Name)
    }
}

// Test `Delete()`
func TestDeletePlace(t *testing.T) {
    place := &Place{
        Name:     "To Be Deleted",
        Type:     "bar",
        Address:  "Delete Address",
        District: "South",
        City:     "DeleteCity",
    }
    testRepo.Create(place)

    err := testRepo.Delete(place.ID)
    if err != nil {
        t.Fatalf("Error deleting place: %v", err)
    }

    deleted, _ := testRepo.GetByID(place.ID)
    if deleted != nil {
        t.Errorf("Delete failed: expected nil, got '%s'", deleted.Name)
    }
}
