package cli

import (
	"github.com/rafaeltiribas/cardapio-uff/internal/delivery"
	"github.com/rafaeltiribas/cardapio-uff/internal/usecase"
)

func runApplication() {
	usecase.PrepareApplication()
	delivery.StartBot()
}

func Execute() {
	runApplication()
}
