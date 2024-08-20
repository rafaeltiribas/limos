package usecase

import (
	"fmt"
	"github.com/rafaeltiribas/cardapio-uff/internal/models"
	"strings"
)

func formatCardapio(refeicao models.Refeicao) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Cardápio do dia %s:\n", refeicao.Date))
	sb.WriteString(fmt.Sprintf("🍽️ %s\n", refeicao.Refeicao))
	sb.WriteString("Prato principal:\n")
	sb.WriteString(fmt.Sprintf("🍗 %s\n", refeicao.Proteina))
	sb.WriteString("Filézinho de frango, contém derivados de soja, trigo e glúten\n")

	sb.WriteString("Acompanhamento:\n")
	sb.WriteString(fmt.Sprintf("🍚 %s\n", refeicao.Acompanhamento))

	sb.WriteString("Guarnição:\n")
	sb.WriteString(fmt.Sprintf("🥕 %s\n", refeicao.Guarnicao))
	sb.WriteString("(Cenoura, cebola, passas, ervilha, óleo de soja, azeite, alho e sal)\n")

	sb.WriteString("Salada:\n")
	sb.WriteString(fmt.Sprintf("🥬 %s\n", refeicao.Salada1))
	sb.WriteString(fmt.Sprintf("🥒 %s\n", refeicao.Salada2))

	sb.WriteString("Sobremesa:\n")
	sb.WriteString(fmt.Sprintf("🍎 %s\n", refeicao.Sobremesa))

	return sb.String()
}
