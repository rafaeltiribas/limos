package cli

import "github.com/rafaeltiribas/cardapio-uff/internal/delivery"

func getCardapioAndSend() {
	delivery.RetrieveMenu()
}

func Execute() {
	getCardapioAndSend()
}
