package usecase

import (
	"github.com/rafaeltiribas/cardapio-uff/internal/models"
)

func CreateMessage(cardapio models.Cardapio) string {
	var message string
	for _, refeicao := range cardapio.Refeicoes {
		message += FormatCardapio(refeicao)
	}

	return message
}
