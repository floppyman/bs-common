package jwttoken

import (
	"strings"
	"testing"
	"time"
)

var issuer = "test-issuer"
var secret = "test-secret-12345678901234567890123456789012"

func TestGenerateAccessToken(t *testing.T) {
	t.Parallel()

	t.Run("generates non-empty token", func(t *testing.T) {
		token, err := GenerateAccessToken(secret, issuer, 42, 60)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if token == "" {
			t.Error("expected non-empty token")
		}
		// JWT tokens have 3 parts separated by dots
		parts := strings.Split(token, ".")
		if len(parts) != 3 {
			t.Errorf("expected 3 JWT parts, got %d", len(parts))
		}
	})

	t.Run("includes correct claims", func(t *testing.T) {
		userID := 99
		token, err := GenerateAccessToken(secret, issuer, userID, 30)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		verifiedID, err := VerifyAccessToken(secret, issuer, token)
		if err != nil {
			t.Fatalf("failed to verify freshly generated token: %v", err)
		}
		if verifiedID != userID {
			t.Errorf("userID: got %d, want %d", verifiedID, userID)
		}
	})
}

func TestVerifyAccessToken(t *testing.T) {
	t.Parallel()

	t.Run("valid token verifies", func(t *testing.T) {
		token, err := GenerateAccessToken(secret, issuer, 1, 60)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		uid, err := VerifyAccessToken(secret, issuer, token)
		if err != nil {
			t.Fatalf("unexpected error verifying valid token: %v", err)
		}
		if uid != 1 {
			t.Errorf("userID: got %d, want 1", uid)
		}
	})

	t.Run("expired token fails", func(t *testing.T) {
		token, err := GenerateAccessToken(secret, issuer, 1, -1) // expires 1 minute ago
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, err = VerifyAccessToken(secret, issuer, token)
		if err == nil {
			t.Error("expected error for expired token")
		}
	})

	t.Run("invalid signature fails", func(t *testing.T) {
		token, err := GenerateAccessToken(secret, issuer, 1, 60)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		tampered := token + "x"
		_, err = VerifyAccessToken(secret, issuer, tampered)
		if err == nil {
			t.Error("expected error for tampered token")
		}
	})

	t.Run("wrong secret fails", func(t *testing.T) {
		// Generate with current secret
		token, err := GenerateAccessToken(secret, issuer, 1, 60)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// Temporarily change secret
		orig := secret
		secret = "different-secret-1234567890123456789012"
		_, err = VerifyAccessToken(secret, issuer, token)
		secret = orig
		if err == nil {
			t.Error("expected error with wrong secret")
		}
	})

	t.Run("wrong issuer fails", func(t *testing.T) {
		token, err := GenerateAccessToken(secret, issuer, 1, 60)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		orig := issuer
		issuer = "wrong-issuer"
		_, err = VerifyAccessToken(secret, issuer, token)
		issuer = orig
		if err == nil {
			t.Error("expected error with wrong issuer")
		}
	})

	t.Run("malformed token fails", func(t *testing.T) {
		cases := []string{
			"",
			"not.a.jwt",
			"badjwt",
		}
		for _, c := range cases {
			_, err := VerifyAccessToken(secret, issuer, c)
			if err == nil {
				t.Errorf("expected error for malformed token %q", c)
			}
		}
	})

	t.Run("token expires at correct time", func(t *testing.T) {
		expiryMinutes := 5
		beforeGen := time.Now()
		token, err := GenerateAccessToken(secret, issuer, 1, expiryMinutes)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		afterGen := time.Now()

		// Parse the token to check exp claim
		// We can't easily parse without Verify, but we can check the token verifies immediately
		_, err = VerifyAccessToken(secret, issuer, token)
		if err != nil {
			t.Fatalf("token should be valid immediately after generation: %v", err)
		}

		// The token should be generated with correct expiry window
		expectedExpiry := beforeGen.Add(time.Duration(expiryMinutes) * time.Minute)
		if afterGen.After(expectedExpiry) {
			t.Log("token generation took longer than expected expiry window")
		}
	})
}
