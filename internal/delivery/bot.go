package delivery

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rafaeltiribas/cardapio-uff/internal/config"
	"github.com/rafaeltiribas/cardapio-uff/internal/security"
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

	limiter := security.NewRateLimiter(config.RequestLimit, config.BanDuration, config.WindowDuration)

	go usecase.SendMenuUpdatesMessages(bot)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID

		if !limiter.IsAllowed(chatID) {
			msg := tgbotapi.NewMessage(chatID, "Você está enviando muitas mensagens. Por favor, espere um momento antes de tentar novamente.")
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = usecase.Command(update.Message.Command(), chatID)
		msg.ParseMode = "MarkdownV2"

		if _, err := bot.Send(msg); err != nil {
			msg.Text = usecase.EscapeMarkdown(config.EmptyMenu)
			msg.ParseMode = "MarkdownV2"
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			log.Println(err)
		}
	}
}
