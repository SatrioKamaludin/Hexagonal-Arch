// Loads environment variables and configurations
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI       string
	DBName         string
	PostgresUser   string
	PostgresPass   string
	PostgresHost   string
	PostgresPort   string
	PostgresDBName string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		MongoURI:       os.Getenv("MONGO_URI"),
		DBName:         os.Getenv("DB_NAME"),
		PostgresUser:   os.Getenv("POSTGRES_USER"),
		PostgresPass:   os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:   os.Getenv("POSTGRES_HOST"),
		PostgresPort:   os.Getenv("POSTGRES_PORT"),
		PostgresDBName: os.Getenv("POSTGRES_DB"),
	}
}
