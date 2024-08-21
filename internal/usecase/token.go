package usecase

import (
	"github.com/joho/godotenv"
	"github.com/rafaeltiribas/cardapio-uff/internal/config"
	"log"
	"os"
	"path/filepath"
)

func RetrieveToken(tokenKey string) string {
	envPath, err := filepath.Abs(config.EnvPath)
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
