package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	RedisURL    string
	JWTSecret   string

	MinIOEndpoint  string
	MinIOAccessKey string
	MinIOSecretKey string
	MinIOBucket    string

	EmailSMTPHost     string
	EmailSMTPPort     int
	EmailSMTPUser     string
	EmailSMTPPassword string
	EmailFrom         string

	AppURL string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/messenger?sslmode=disable"),
		RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key-change-me"),

		MinIOEndpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
		MinIOAccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
		MinIOSecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
		MinIOBucket:    getEnv("MINIO_BUCKET", "messenger"),

		EmailSMTPHost:     getEnv("EMAIL_SMTP_HOST", "smtp.gmail.com"),
		EmailSMTPPort:     getEnvAsInt("EMAIL_SMTP_PORT", 587),
		EmailSMTPUser:     getEnv("EMAIL_SMTP_USER", ""),
		EmailSMTPPassword: getEnv("EMAIL_SMTP_PASSWORD", ""),
		EmailFrom:         getEnv("EMAIL_FROM", "noreply@messenger.com"),

		AppURL: getEnv("APP_URL", "http://localhost:3000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	strValue := getEnv(key, "")
	if strValue == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(strValue)
	if err != nil {
		log.Printf("Invalid integer for %s: %s, using default %d", key, strValue, defaultValue)
		return defaultValue
	}
	return value
}
