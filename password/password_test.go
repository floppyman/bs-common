package password

import (
	"strings"
	"testing"
)

func TestGenerateRandomPassword(t *testing.T) {
	t.Parallel()

	t.Run("generates password of requested length", func(t *testing.T) {
		for _, length := range []int{8, 16, 32, 64} {
			pw, err := GenerateRandomPassword(length)
			if err != nil {
				t.Fatalf("length %d: unexpected error: %v", length, err)
			}
			if len(pw) != length {
				t.Errorf("length %d: got %d, want %d", length, len(pw), length)
			}
		}
	})

	t.Run("password contains only allowed characters", func(t *testing.T) {
		pw, err := GenerateRandomPassword(100)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		allowed := AlphabetCharset
		for _, c := range pw {
			if !strings.ContainsRune(allowed, c) {
				t.Errorf("disallowed character %q in password", c)
			}
		}
	})

	t.Run("returns empty string for length zero", func(t *testing.T) {
		pw, err := GenerateRandomPassword(0)
		if err != nil {
			t.Fatalf("unexpected error for length 0: %v", err)
		}
		if pw != "" {
			t.Errorf("expected empty string for length 0, got %q", pw)
		}
	})
}

func TestHashPassword(t *testing.T) {
	t.Parallel()

	t.Run("produces argon2id formatted hash", func(t *testing.T) {
		hash := HashPassword("mysecret")
		if !strings.HasPrefix(hash, "$argon2id$v=19$") {
			t.Errorf("unexpected hash prefix: %s", hash)
		}
		parts := strings.Split(hash, "$")
		if len(parts) != 6 {
			t.Errorf("expected 6 parts, got %d: %s", len(parts), hash)
		}
	})

	t.Run("different hashes for same password", func(t *testing.T) {
		hash1 := HashPassword("mysecret")
		hash2 := HashPassword("mysecret")
		if hash1 == hash2 {
			t.Error("expected different hashes due to random salt, got identical")
		}
	})
}

func TestVerifyPassword(t *testing.T) {
	t.Parallel()

	t.Run("valid password verifies", func(t *testing.T) {
		password := "correcthorsebatterystaple"
		hash := HashPassword(password)
		if !VerifyPassword(password, hash) {
			t.Error("expected valid password to verify")
		}
	})

	t.Run("invalid password fails", func(t *testing.T) {
		hash := HashPassword("correcthorsebatterystaple")
		if VerifyPassword("wrongpassword", hash) {
			t.Error("expected invalid password to fail verification")
		}
	})

	t.Run("malformed hash returns false", func(t *testing.T) {
		cases := []string{
			"",
			"notahash",
			"$argon2i$v=19$m=47104,t=1,p=1$salt$hash",
			"$argon2id$v=18$m=47104,t=1,p=1$salt$hash",
			"$argon2id$v=19$m=47104,t=1,p=1$salt",
			"$argon2id$v=19$m=bad,t=1,p=1$salt$hash",
		}
		for _, c := range cases {
			if VerifyPassword("password", c) {
				t.Errorf("expected false for malformed hash %q", c)
			}
		}
	})

	t.Run("tampered hash returns false", func(t *testing.T) {
		password := "mysecret"
		hash := HashPassword(password)
		tampered := hash + "extra"
		if VerifyPassword(password, tampered) {
			t.Error("expected tampered hash to fail")
		}
	})
}

func TestValidatePassword(t *testing.T) {
	t.Parallel()

	t.Run("accepts password at minimum length", func(t *testing.T) {
		if err := ValidatePassword(strings.Repeat("a", MinPasswordLength)); err != nil {
			t.Errorf("expected no error for %d char password, got %v", MinPasswordLength, err)
		}
	})

	t.Run("rejects password below minimum length", func(t *testing.T) {
		if err := ValidatePassword(strings.Repeat("a", MinPasswordLength-1)); err == nil {
			t.Errorf("expected error for %d char password", MinPasswordLength-1)
		}
	})
}
