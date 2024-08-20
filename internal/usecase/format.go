package usecase

import (
	"fmt"
	"github.com/rafaeltiribas/cardapio-uff/internal/models"
	"strings"
)

func formatCardapio(cardapio models.Cardapio) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Cardápio do dia %s:\n", node.Date))
	sb.WriteString(fmt.Sprintf("🍽️ %s\n", node.Refeicao))
	sb.WriteString("Prato principal:\n")
	sb.WriteString(fmt.Sprintf("🍗 %s\n", node.Proteina))
	sb.WriteString("Filézinho de frango, contém derivados de soja, trigo e glúten\n")

	sb.WriteString("Acompanhamento:\n")
	sb.WriteString(fmt.Sprintf("🍚 %s\n", node.Acompanhamento))

	sb.WriteString("Guarnição:\n")
	sb.WriteString(fmt.Sprintf("🥕 %s\n", node.Guarnicao))
	sb.WriteString("(Cenoura, cebola, passas, ervilha, óleo de soja, azeite, alho e sal)\n")

	sb.WriteString("Salada:\n")
	sb.WriteString(fmt.Sprintf("🥬 %s\n", node.Salada1))
	sb.WriteString(fmt.Sprintf("🥒 %s\n", node.Salada2))

	sb.WriteString("Sobremesa:\n")
	sb.WriteString(fmt.Sprintf("🍎 %s\n", node.Sobremesa))

	return sb.String()
}
