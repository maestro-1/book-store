package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	// JWTSecret              string
	// JWTExpirationInSeconds int64
}

func(s *Config) createDatabaseURI() string {
	databaseUri := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", s.DBUser, s.DBPassword, s.DBAddress, s.DBName);
	fmt.Println(databaseUri)
	return databaseUri;
}



var Envs = initConfig()
var DatabaseUri = Envs.createDatabaseURI()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:             getEnv("HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "maestro"),
		DBPassword:             getEnv("DB_PASSWORD", "maestro"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
		DBName:                 getEnv("DB_NAME", "book-store"),
	}
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}