package usecase

import (
	"fmt"
	"github.com/rafaeltiribas/cardapio-uff/internal/models"
	"strings"
)

func FormatCardapio(refeicao models.Refeicao) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("📅 Cardápio do dia: %s\n", refeicao.Date))
	sb.WriteString(fmt.Sprintf("🍽️ Refeição: %s\n\n", refeicao.Refeicao))

	sb.WriteString("🥩 *Prato Principal*\n")
	sb.WriteString(fmt.Sprintf("\\- %s\n", EscapeMarkdown(refeicao.Proteina))) // Aqui você escapa os textos
	sb.WriteString("🔍 Contém derivados de soja, trigo e glúten\n\n")

	sb.WriteString("🍚 *Acompanhamento*\n")
	sb.WriteString(fmt.Sprintf("\\- %s\n\n", EscapeMarkdown(refeicao.Acompanhamento)))

	sb.WriteString("🍴 *Guarnição*\n")
	sb.WriteString(fmt.Sprintf("\\- %s\n\n", EscapeMarkdown(refeicao.Guarnicao)))

	sb.WriteString("🥗 *Salada*\n")
	sb.WriteString(fmt.Sprintf("\\- %s\n", EscapeMarkdown(refeicao.Salada1)))
	sb.WriteString(fmt.Sprintf("\\- %s\n\n", EscapeMarkdown(refeicao.Salada2)))

	sb.WriteString("🍫 *Sobremesa*\n")
	sb.WriteString(fmt.Sprintf("\\- %s\n\n", EscapeMarkdown(refeicao.Sobremesa)))

	return sb.String()
}
