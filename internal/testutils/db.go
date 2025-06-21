package testutils

import (
	"database/sql"
	"testing"

	"example.com/rest-api/db"
	_ "github.com/mattn/go-sqlite3"
)

// SetupTestDB initializes an in-memory SQLite database and creates schema.
func SetupTestDB(t *testing.T) func() {
	var err error
	db.DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	createUsers := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        );`
	if _, err = db.DB.Exec(createUsers); err != nil {
		t.Fatalf("create users table: %v", err)
	}
	createEvents := `
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INTEGER,
            FOREIGN KEY(user_id) REFERENCES users(id)
        );`
	if _, err = db.DB.Exec(createEvents); err != nil {
		t.Fatalf("create events table: %v", err)
	}
	createRegs := `
        CREATE TABLE IF NOT EXISTS registrations (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            event_id INTEGER,
            user_id INTEGER,
            FOREIGN KEY(event_id) REFERENCES events(id),
            FOREIGN KEY(user_id) REFERENCES users(id)
        );`
	if _, err = db.DB.Exec(createRegs); err != nil {
		t.Fatalf("create registrations table: %v", err)
	}
	return func() { db.DB.Close() }
}
