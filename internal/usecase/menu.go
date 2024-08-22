package usecase

func Menu() string {
	var mensagem string

	if CheckFile() {
		mensagem = GetFileContent()
	} else {
		cardapio := RetrieveMenu()
		mensagem = CreateMessage(cardapio)
		CreateFile(mensagem)
	}

	return mensagem
}

func MenuHasChanged() bool {
	return true
}

func LatestMenu() string {
	var mensagem string

	cardapio := RetrieveMenu()
	mensagem = CreateMessage(cardapio)
	CreateLatestMenuFile(mensagem)

	return mensagem
}

func GetLatestMenu() string {
	var mensagem string

	cardapio := RetrieveMenu()
	mensagem = CreateMessage(cardapio)

	return mensagem
}
