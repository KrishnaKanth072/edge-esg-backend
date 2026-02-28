package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	RedisURL      string
	KeycloakURL   string
	KeycloakRealm string
	ServerPort    string
	EncryptionKey string
	Environment   string
	TLSEnabled    bool
	TLSCertFile   string
	TLSKeyFile    string
}

func Load() (*Config, error) {
	_ = godotenv.Load() // Ignore error if .env file doesn't exist

	config := &Config{
		DatabaseURL:   getEnvRequired("DATABASE_URL"),
		RedisURL:      getEnvRequired("REDIS_URL"),
		KeycloakURL:   getEnvRequired("KEYCLOAK_URL"),
		KeycloakRealm: getEnvRequired("KEYCLOAK_REALM"),
		ServerPort:    getEnv("SERVER_PORT", "8000"),
		EncryptionKey: getEnvRequired("ENCRYPTION_KEY"),
		Environment:   getEnv("ENVIRONMENT", "development"),
		TLSEnabled:    getEnv("TLS_ENABLED", "false") == "true",
		TLSCertFile:   getEnv("TLS_CERT_FILE", ""),
		TLSKeyFile:    getEnv("TLS_KEY_FILE", ""),
	}

	// Validate encryption key length
	if len(config.EncryptionKey) != 64 {
		return nil, fmt.Errorf("ENCRYPTION_KEY must be 64 hex characters (32 bytes)")
	}

	// Enforce TLS in production
	if config.Environment == "production" && !config.TLSEnabled {
		return nil, fmt.Errorf("TLS must be enabled in production environment")
	}

	return config, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Required environment variable %s is not set", key))
	}
	return value
}
