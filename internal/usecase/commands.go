package usecase

import (
	"fmt"
	"github.com/rafaeltiribas/cardapio-uff/internal/repository"
)

func Command(command string, chatID int64) string {
	switch command {
	case "start":
		repository.NewUser(chatID)

		return EscapeMarkdown("Olá! Eu sou o bot do Bandeijão da UFF. Aqui estão algumas coisas que posso fazer:\n\n" +
			"🍽️ /cardapio - Veja o cardápio do dia\n" +
			"🕒 /horarios - Confira os horários de funcionamento\n" +
			"ℹ️ /sobre - Saiba mais sobre o bot\n" +
			"📞 /contato - Entre em contato comigo\n" +
			"❓ /ajuda - Lista todos os comandos que eu entendo\n\n" +
			"Digite um comando para começar!")
	case "ajuda":
		return EscapeMarkdown("Aqui estão os comandos que você pode usar:\n\n" +
			"🍽️ /cardapio - Veja o cardápio do dia\n" +
			"🕒 /horarios - Confira os horários de funcionamento\n" +
			"ℹ️ /sobre - Saiba mais sobre o bot\n" +
			"📞 /contato - Entre em contato comigo\n" +
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
			"🍽️ Almoço: 11h15 - 14h30 (Gragoatá) e 12h00 - 14h00 (Praia Vermelha)\n" +
			"🍲 Jantar: 17h00 - 19h00\n\n" +
			"Qualquer dúvida, é só usar o comando /ajuda!")
	case "sobre":
		return EscapeMarkdown("ℹ️ Sobre o Bot\n\n" +
			"Este é um projeto desenvolvido por um estudante da UFF para tornar o acesso ao cardápio do bandeijão " +
			"mais prático. Consulte a qualquer momento o cardápio atualizado do bandeijão ou espere receber em primeira mão o menu do dia " +
			"assim que sair (antes mesmo de postarem no Instagram!).\n\n" +
			"Desenvolvido por: @techtiribas\n" +
			"Fase: Testes\n\n" +
			"O projeto está em fase de testes e pode apresentar alguns erros. Agradeço o feedback!")
	case "contato":
		return EscapeMarkdown("📞 Contato\n\n" +
			"Este é um serviço desenvolvido de estudante para estudante. Se você tiver alguma dúvida, sugestão ou encontrar algum problema, entre em contato pelos seguintes meios:\n\n" +
			"🐦 Twitter: @techtiribas\n" +
			"🩶 Discord: @techtiribas\n" +
			"📧 Email: rafaeltiribas@outlook.com\n" +
			"🌐 LinkedIn: https://www.linkedin.com/in/rafaeltiribas/\n\n" +
			"Agradeço o feedback e sinta-se à vontade para contribuir com o projeto!")
	default:
		return EscapeMarkdown("😕 Oops! Eu não reconheço esse comando.\n\n" +
			"Tente usar um dos comandos que eu conheço:\n" +
			"🍽️ /cardapio - Veja o cardápio do dia\n" +
			"🕒 /horarios - Confira os horários de funcionamento\n" +
			"ℹ️ /sobre - Saiba mais sobre o bot\n" +
			"❓ /ajuda - Lista todos os comandos disponíveis\n\n" +
			"Se precisar de ajuda, use /ajuda para ver a lista completa.")
	}
}
