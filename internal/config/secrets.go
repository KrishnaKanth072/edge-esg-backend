package config

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

// GetBankKey retrieves bank-specific encryption key from environment
// Keys must be set as: BANK_KEY_<BANK_ID>=<32-byte-hex-key>
func GetBankKey(bankID string) (string, error) {
	envKey := fmt.Sprintf("BANK_KEY_%s", bankID)
	key := os.Getenv(envKey)

	if key == "" {
		return "", fmt.Errorf("encryption key not found for bank: %s. Set %s environment variable", bankID, envKey)
	}

	if len(key) != 64 { // 32 bytes = 64 hex characters
		return "", fmt.Errorf("invalid key length for bank %s. Must be 64 hex characters (32 bytes)", bankID)
	}

	return key, nil
}

// GenerateSecureKey generates a cryptographically secure 32-byte key
// Use this to generate keys for each bank, then store in secrets manager
func GenerateSecureKey() (string, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

// ValidateBankID validates bank ID format
func ValidateBankID(bankID string) bool {
	// Bank IDs should be alphanumeric and between 8-64 characters
	if len(bankID) < 8 || len(bankID) > 64 {
		return false
	}

	for _, char := range bankID {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '-' || char == '_') {
			return false
		}
	}

	return true
}
