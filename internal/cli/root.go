package cli

import "github.com/rafaeltiribas/cardapio-uff/internal/delivery"

func runApplication() {
	delivery.StartBot()
}

func Execute() {
	runApplication()
}
