package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rafaeltiribas/cardapio-uff/internal/usecase"
	"log"
)

func StartBot() {
	token := usecase.RetrieveToken("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		msg.Text = usecase.Command(update.Message.Command())
		
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
