package usecase

import (
	"github.com/rafaeltiribas/cardapio-uff/internal/config"
	"log"
	"os"
	"time"
)

func CheckFile() bool {
	if _, err := os.Stat(config.TempFilePath); err == nil {
		return true
	}

	return false
}

func CheckWeekday() bool {
	now := time.Now()

	if now.Weekday() >= time.Monday && now.Weekday() <= time.Friday {
		return true
	}

	return false
}

func CheckWorkingHour() bool {
	now := time.Now()

	if now.Hour() >= 8 && now.Hour() <= 17 {
		return true
	}

	return false
}

func CheckMenuChanges() bool {
	old := GetOldMenuFileContent()
	latest := GetLatestMenu()

	if latest != old && latest != "" {
		log.Println("Menu updated!")
		return true
	}

	log.Println("Same menu or new menu is empty.")
	return false

}
