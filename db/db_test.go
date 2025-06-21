package db

import "testing"

func TestInitDB(t *testing.T) {
	InitDB()
	if DB == nil {
		t.Fatal("DB should not be nil after init")
	}
	if err := DB.Ping(); err != nil {
		t.Fatalf("ping error: %v", err)
	}
	DB.Close()
}
