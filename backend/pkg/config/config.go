// Package config loads application configuration from environment variables.
package config

import "os"

// Config holds all runtime configuration.
type Config struct {
	Env      string // "development" | "production"
	HTTPPort string // port for the REST API server, default "8080"
	BotToken string // Telegram bot token (required for bot binary)
}

// Load reads configuration from environment variables.
// Missing optional fields fall back to defaults.
func Load() *Config {
	cfg := &Config{
		Env:      getEnv("ENV", "development"),
		HTTPPort: getEnv("HTTP_PORT", "8080"),
		BotToken: os.Getenv("BOT_TOKEN"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
