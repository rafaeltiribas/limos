package delivery

import (
	"github.com/rafaeltiribas/cardapio-uff/internal/usecase"
	"log"
	"os"
)

func Menu() string {
	var mensagem string

	if usecase.CheckFile() {
		response, err := os.ReadFile("../../internal/usecase/temp/tempcardapio")
		if err != nil {
			log.Fatal(err)
		}
		mensagem = string(response)
	} else {
		cardapio := retrieveMenu()
		mensagem = usecase.CreateMessage(cardapio)
		usecase.CreateFile(mensagem)
	}

	return mensagem
}
