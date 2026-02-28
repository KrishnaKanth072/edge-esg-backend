package config

import "github.com/google/uuid"

// Per-bank encryption keys (RLS isolation)
var BankEncryptionKeys = map[string]string{
	"hdfc-bank-uuid":  "hdfc-aes256-key-32bytes-secret",
	"icici-bank-uuid": "icici-aes256-key-32bytes-secret",
	"sbi-bank-uuid":   "sbi-aes256-key-32bytes-secret",
}

func GetBankKey(bankID string) string {
	if key, exists := BankEncryptionKeys[bankID]; exists {
		return key
	}
	return "default-aes256-key-32bytes-secret"
}

func ValidateBankID(bankID string) bool {
	_, err := uuid.Parse(bankID)
	return err == nil
}
