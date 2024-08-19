package main

import (
	"encoding/xml"
	"fmt"
	"github.com/rafaeltiribas/cardapio-uff/internal/models"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.restaurante.uff.br/cardapiomobile.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var cardapio models.Cardapio
	err = xml.Unmarshal(body, &cardapio)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, refeicao := range cardapio.Refeicoes {
		fmt.Println(refeicao)
	}

}
