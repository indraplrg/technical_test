package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string
	AppEnv       string
	AppPort      string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	DBTimeZone   string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

func Load(paths ...string) *Config {
	_ = godotenv.Load(paths...)

	return &Config{
		AppName:      getEnv("APP_NAME", "student-management"),
		AppEnv:       getEnv("APP_ENV", "development"),
		AppPort:      getEnv("APP_PORT", "8080"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "postgres"),
		DBName:       getEnv("DB_NAME", "student_db"),
		DBSSLMode:    getEnv("DB_SSLMODE", "disable"),
		DBTimeZone:   getEnv("DB_TIMEZONE", "UTC"),
		ReadTimeout:  getEnvInt("READ_TIMEOUT", 10),
		WriteTimeout: getEnvInt("WRITE_TIMEOUT", 10),
		IdleTimeout:  getEnvInt("IDLE_TIMEOUT", 60),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
