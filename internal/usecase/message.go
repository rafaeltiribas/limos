package usecase

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rafaeltiribas/cardapio-uff/internal/models"
	"github.com/rafaeltiribas/cardapio-uff/internal/repository"
	"log"
	"time"
)

func CreateMessage(cardapio models.Cardapio) string {
	var message string
	for _, refeicao := range cardapio.Refeicoes {
		message += FormatCardapio(refeicao)
	}

	return message
}

func SendMenuUpdatesMessages(bot *tgbotapi.BotAPI) {

	ticker := time.NewTicker(1 * time.Hour)

	for range ticker.C {
		if CheckWeekday() {
			if CheckWorkingHour() {
				if CheckMenuChanges() {
					SendMessageToAllUsers(bot)
				}
			}
		}
	}
}

func SendMessageToAllUsers(bot *tgbotapi.BotAPI) {
	chatIDs, err := repository.GetAllChatIDs()
	if err != nil {
		log.Println("Error getting all IDs: %v", err)
		return
	}

	newMenu := LatestMenu()

	for _, chatID := range chatIDs {
		msg := tgbotapi.NewMessage(chatID, newMenu)
		msg.ParseMode = "MarkdownV2"

		if _, err := bot.Send(msg); err != nil {
			log.Println("Error sending update message: %v", err)
		}
	}
}
