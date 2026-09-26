package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	argon2Memory      = 47104
	argon2Iterations  = 1
	argon2Parallelism = 1
	argon2KeyLen      = 32
	saltLen           = 16

	MinPasswordLength = 16

	AlphabetCharset   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-"
)

// ValidatePassword returns an error if the password does not meet complexity requirements.
func ValidatePassword(password string) error {
	if len(password) < MinPasswordLength {
		return fmt.Errorf("password must be at least %d characters", MinPasswordLength)
	}
	return nil
}

// GenerateRandomPassword creates a cryptographically secure random alphanumeric password of the given length.
func GenerateRandomPassword(length int) (string, error) {
	password := make([]byte, length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(AlphabetCharset))))
		if err != nil {
			return "", err
		}
		password[i] = AlphabetCharset[n.Int64()]
	}
	return string(password), nil
}

// HashPassword hashes the password using argon2id with a randomly generated salt and returns a compact encoded string.
func HashPassword(password string) string {
	salt := make([]byte, saltLen)
	_, _ = rand.Read(salt)
	saltStr := base64.RawURLEncoding.EncodeToString(salt)
	hash := argon2.IDKey([]byte(password), []byte(saltStr), argon2Iterations, argon2Memory, argon2Parallelism, argon2KeyLen)
	return fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		argon2Memory, argon2Iterations, argon2Parallelism,
		saltStr, base64.RawURLEncoding.EncodeToString(hash))
}

// VerifyPassword verifies a plaintext password against a stored argon2id hash.
func VerifyPassword(password, encodedHash string) bool {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false
	}
	if parts[1] != "argon2id" {
		return false
	}
	if parts[2] != "v=19" {
		return false
	}

	var memory, iterations uint32
	var parallelism uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return false
	}

	salt := []byte(parts[4])

	expectedHash, err := base64.RawURLEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, uint32(len(expectedHash)))
	return subtle.ConstantTimeCompare(hash, expectedHash) == 1
}
