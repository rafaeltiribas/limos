package usecase

func Command(command string) string {
	switch command {
	case "start":
		return "Olá! Eu sou o bot do RU da UFF. Aqui estão algumas coisas que posso fazer:\n\n" +
			"🍽️ /cardapio - Veja o cardápio do dia\n" +
			"🕒 /horarios - Confira os horários de funcionamento\n" +
			"❓ /ajuda - Lista todos os comandos que eu entendo\n\n" +
			"Digite um comando para começar!"
		//adicionar o id do usuário em um bd para ter uma contagem de usuários do bot.
	case "ajuda":
		return "Aqui estão os comandos que você pode usar:\n\n" +
			"🍽️ /cardapio - Veja o cardápio do dia\n" +
			"🕒 /horarios - Confira os horários de funcionamento\n" +
			"🏠 /start - Mostra a mensagem de boas-vindas e os comandos principais\n\n" +
			"Se precisar de mais alguma coisa, é só perguntar!"
	case "cardapio":
		return Menu()
	case "horarios":
		return "🕒 Horários de funcionamento do RU da UFF:\n\n" +
			"🍽️ Almoço: 11h30 - 14h00\n" +
			"🍲 Jantar: 17h30 - 19h30\n\n" +
			"Qualquer dúvida, é só usar o comando /ajuda!"
	default:
		return "😕 Oops! Eu não reconheço esse comando.\n\n" +
			"Tente usar um dos comandos que eu conheço:\n" +
			"🍽️ /cardapio - Veja o cardápio do dia\n" +
			"🕒 /horarios - Confira os horários de funcionamento\n" +
			"❓ /ajuda - Lista todos os comandos disponíveis\n\n" +
			"Se precisar de ajuda, use /ajuda para ver a lista completa."
	}
}
