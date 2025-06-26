package utils

import "testing"

func TestGenerateAndVerifyToken(t *testing.T) {
	token, err := GenerateToken("user@example.com", 1)
	if err != nil {
		t.Fatalf("generate token: %v", err)
	}
	id, err := VerifyToken(token)
	if err != nil {
		t.Fatalf("verify token: %v", err)
	}
	if id != 1 {
		t.Errorf("expected id 1, got %d", id)
	}
}

func TestVerifyTokenFail(t *testing.T) {
	if _, err := VerifyToken("badtoken"); err == nil {
		t.Fatal("expected error for invalid token")
	}
}
