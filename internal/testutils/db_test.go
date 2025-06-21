package testutils

import (
	"example.com/rest-api/db"
	"testing"
)

func TestSetupTestDB(t *testing.T) {
	cleanup := SetupTestDB(t)
	if db.DB == nil {
		t.Fatal("DB should not be nil")
	}
	if err := db.DB.Ping(); err != nil {
		t.Fatalf("ping error: %v", err)
	}
	cleanup()
}
