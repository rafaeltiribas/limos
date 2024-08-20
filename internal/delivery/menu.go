package delivery

import (
	"fmt"
	"github.com/rafaeltiribas/cardapio-uff/internal/usecase"
)

func Menu() {
	cardapio := retrieveMenu()
	mensagem := usecase.CreateMessage(cardapio)
	fmt.Println(mensagem)
}
