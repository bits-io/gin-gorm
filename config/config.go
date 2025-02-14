package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	JWTSecret   string
	DBUsername  string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
}

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("error while load .env file")
	}
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		JWTSecret:   os.Getenv("JWT_SECRET"),
		DBUsername:  os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBName:      os.Getenv("DB_NAME"),
	}
}