package models

import (
	"testing"

	"example.com/rest-api/internal/testutils"
)

func TestUserSaveAndValidate(t *testing.T) {
	cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	user := User{Email: "test@example.com", Password: "secret"}
	if err := user.Save(); err != nil {
		t.Fatalf("save user: %v", err)
	}

	u := User{Email: "test@example.com", Password: "secret"}
	if err := u.ValidateCredentials(); err != nil {
		t.Fatalf("validate credentials: %v", err)
	}
	if u.ID == 0 {
		t.Error("expected user ID to be set")
	}
}
