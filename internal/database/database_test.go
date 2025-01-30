package database

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const testDB = "postgres://myuser:mypassword@localhost:5432/ruangketiga_test?sslmode=disable"

func TestConnect(t *testing.T) {
    os.Setenv("DB_HOST", "localhost")
    os.Setenv("DB_PORT", "5432")
    os.Setenv("DB_USER", "myuser")
    os.Setenv("DB_PASSWORD", "mypassword")
    os.Setenv("DB_NAME", "ruangketiga_test")

    db, err := Connect()
    if err != nil {
        t.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        t.Errorf("Database connection failed: %v", err)
    }
}

func TestInvalidDBConnection(t *testing.T) {
    os.Setenv("DB_USER", "invalid_user")
    os.Setenv("DB_PASSWORD", "wrong_password")

    _, err := Connect()
    if err == nil {
        t.Errorf("Expected an error for invalid database credentials, but got none")
    }
}
