package models

type Refeicao struct {
	Date           string `xml:"no-name"`
	Refeicao       string `xml:"Tipo-de-refei-o"`
	Proteina       string `xml:"Prato-principal"`
	Guarnicao      string `xml:"Guarni-o"`
	Acompanhamento string `xml:"Acompanhamentos"`
	Salada1        string `xml:"Salada-1"`
	Salada2        string `xml:"Salada-2"`
	Sobremesa      string `xml:"Sobremesa"`
}
