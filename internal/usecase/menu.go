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
