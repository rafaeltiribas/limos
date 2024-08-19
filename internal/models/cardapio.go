package models

import "encoding/xml"

type Cardapio struct {
	ChaveRefeicao xml.Name   `xml:"nodes"`
	Refeicoes     []Refeicao `xml:"node"`
}
