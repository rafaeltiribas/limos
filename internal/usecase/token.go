package usecase

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func RetrieveToken(tokenKey string) string {
	envPath, err := filepath.Abs("../../.env")
	if err != nil {
		log.Fatalf("Failed to get absolute path of .env: %v", err)
	}

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	token := os.Getenv(tokenKey)

	return token
}
