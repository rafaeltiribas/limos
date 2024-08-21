package usecase

import (
	"encoding/xml"
	"fmt"
	"github.com/rafaeltiribas/cardapio-uff/internal/models"
	"io/ioutil"
	"net/http"
)

func RetrieveMenu() models.Cardapio {
	resp, err := http.Get("https://www.restaurante.uff.br/cardapiomobile.xml")
	if err != nil {
		fmt.Println(err)
		return models.Cardapio{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return models.Cardapio{}
	}

	var cardapio models.Cardapio
	err = xml.Unmarshal(body, &cardapio)
	if err != nil {
		fmt.Println(err)
		return models.Cardapio{}
	}

	return cardapio
}
