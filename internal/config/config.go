package config

import (
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
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://edge_admin:edge_secure_2024@localhost:5432/edge_esg"),
		RedisURL:      getEnv("REDIS_URL", "redis://:edge_redis_2024@localhost:6379/0"),
		KeycloakURL:   getEnv("KEYCLOAK_URL", "http://localhost:8080"),
		KeycloakRealm: getEnv("KEYCLOAK_REALM", "zebbank"),
		ServerPort:    getEnv("SERVER_PORT", "8000"),
		EncryptionKey: getEnv("ENCRYPTION_KEY", "zebbank-edge-2024-aes256-secret-key"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
