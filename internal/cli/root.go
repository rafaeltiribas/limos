package cli

import "github.com/rafaeltiribas/cardapio-uff/internal/delivery"

func getCardapioAndSend() {
	delivery.Menu()
}

func Execute() {
	getCardapioAndSend()
}
