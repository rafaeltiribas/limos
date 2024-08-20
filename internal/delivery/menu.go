package delivery

import (
	"fmt"
)

func Menu() {
	cardapio := retrieveMenu()
	for _, refeicao := range cardapio.Refeicoes {
		fmt.Println(formatCardapio(refeicao))
	}
}
