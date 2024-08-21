package usecase

import (
	"fmt"
	"github.com/rafaeltiribas/cardapio-uff/internal/repository"
)

func Command(command string, chatID int64) string {
	switch command {
	case "start":
		repository.NewUser(chatID)

		return EscapeMarkdown("Olá! Eu sou o bot do RU da UFF. Aqui estão algumas coisas que posso fazer:\n\n" +
			"🍽️ /cardapio - Veja o cardápio do dia\n" +
			"🕒 /horarios - Confira os horários de funcionamento\n" +
			"❓ /ajuda - Lista todos os comandos que eu entendo\n\n" +
			"Digite um comando para começar!")
	case "ajuda":
		return EscapeMarkdown("Aqui estão os comandos que você pode usar:\n\n" +
			"🍽️ /cardapio - Veja o cardápio do dia\n" +
			"🕒 /horarios - Confira os horários de funcionamento\n" +
			"🏠 /start - Mostra a mensagem de boas-vindas e os comandos principais\n\n" +
			"Se precisar de mais alguma coisa, é só perguntar!")
	case "cardapio":
		err := repository.UpdateLastUseDate(chatID)
		if err != nil {
			fmt.Println("Error updating Last Use Date")
		}
		return Menu()
	case "horarios":
		return EscapeMarkdown("🕒 Horários de funcionamento do RU da UFF:\n\n" +
			"🍽️ Almoço: 11h30 - 14h00\n" +
			"🍲 Jantar: 17h30 - 19h30\n\n" +
			"Qualquer dúvida, é só usar o comando /ajuda!")
	default:
		return EscapeMarkdown("😕 Oops! Eu não reconheço esse comando.\n\n" +
			"Tente usar um dos comandos que eu conheço:\n" +
			"🍽️ /cardapio - Veja o cardápio do dia\n" +
			"🕒 /horarios - Confira os horários de funcionamento\n" +
			"❓ /ajuda - Lista todos os comandos disponíveis\n\n" +
			"Se precisar de ajuda, use /ajuda para ver a lista completa.")
	}
}
