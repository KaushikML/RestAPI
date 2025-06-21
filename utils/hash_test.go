package utils

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
	hash, err := HashPassword("secret")
	if err != nil {
		t.Fatalf("hash error: %v", err)
	}
	if !CheckPasswordHash("secret", hash) {
		t.Error("expected password to match")
	}
	if CheckPasswordHash("wrong", hash) {
		t.Error("expected mismatch with wrong password")
	}
}
