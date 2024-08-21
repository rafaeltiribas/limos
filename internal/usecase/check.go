package usecase

import (
	"github.com/rafaeltiribas/cardapio-uff/internal/config"
	"os"
)

func CheckFile() bool {
	if _, err := os.Stat(config.TempFilePath); err == nil {
		return true
	} else {
		return false
	}
}
