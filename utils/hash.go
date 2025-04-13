package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

// GenerateUUID generates a new UUID as string
func GenerateUUID() string {
	return uuid.NewString()
}

// SHA256 hash returns a hex-encoded SHA-256 hash of the input-string
func SHA256Hash(input string) string {
	hash := sha256.Sum256([]byte(input))

	return hex.EncodeToString(hash[:])
}

func HashFileSHA256(path string) (string, error) {

	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hasher := sha256.New()

	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)),nil
}
