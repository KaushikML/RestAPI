package models

import (
	"testing"
	"time"

	"example.com/rest-api/internal/testutils"
)

func TestEventCRUD(t *testing.T) {
	cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	user := User{Email: "e@example.com", Password: "secret"}
	if err := user.Save(); err != nil {
		t.Fatalf("save user: %v", err)
	}
	u := User{Email: "e@example.com", Password: "secret"}
	if err := u.ValidateCredentials(); err != nil {
		t.Fatalf("validate: %v", err)
	}

	event := Event{Name: "Party", Description: "Desc", Location: "Loc", DateTime: time.Now(), UserID: u.ID}
	if err := event.Save(); err != nil {
		t.Fatalf("save event: %v", err)
	}

	fetched, err := GetEventByID(event.ID)
	if err != nil {
		t.Fatalf("get event: %v", err)
	}

	fetched.Name = "Updated"
	if err := fetched.Update(); err != nil {
		t.Fatalf("update event: %v", err)
	}

	updated, err := GetEventByID(event.ID)
	if err != nil {
		t.Fatalf("get updated: %v", err)
	}
	if updated.Name != "Updated" {
		t.Errorf("expected Updated, got %s", updated.Name)
	}

	if err := updated.Delete(); err != nil {
		t.Fatalf("delete event: %v", err)
	}

	// Register and cancel
	event = Event{Name: "Conf", Description: "Talk", Location: "Hall", DateTime: time.Now(), UserID: u.ID}
	if err := event.Save(); err != nil {
		t.Fatalf("save event2: %v", err)
	}
	if err := event.Register(u.ID); err != nil {
		t.Fatalf("register: %v", err)
	}
	if err := event.CancelRegistration(u.ID); err != nil {
		t.Fatalf("cancel: %v", err)
	}

	// Get all events
	events, err := GetAllEvents()
	if err != nil {
		t.Fatalf("get all: %v", err)
	}
	if len(events) == 0 {
		t.Error("expected events list not empty")
	}
}
