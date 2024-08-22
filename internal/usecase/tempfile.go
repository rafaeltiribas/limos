package usecase

import (
	"fmt"
	"github.com/rafaeltiribas/cardapio-uff/internal/config"
	"log"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateFile(msg string) {

	file, err := os.Create(config.TempFilePath)
	check(err)

	_, err = file.WriteString(msg)
	check(err)

	file.Close()

	go func() {
		fmt.Println("Timer On...")
		timer := time.NewTimer(config.ReloadRate)

		<-timer.C
		DeleteFile()
	}()

}

func DeleteFile() {
	err := os.Remove(config.TempFilePath)
	check(err)

}

func GetFileContent() string {
	response, err := os.ReadFile(config.TempFilePath)
	if err != nil {
		log.Fatal(err)
	}
	mensagem := string(response)

	return mensagem
}

func GetOldMenuFileContent() string {
	response, err := os.ReadFile(config.LastFilePath)
	if err != nil {
		log.Fatal(err)
	}
	mensagem := string(response)

	return mensagem
}

func CreateLatestMenuFile(msg string) {

	file, err := os.Create(config.LastFilePath)
	check(err)

	_, err = file.WriteString(msg)
	check(err)

	file.Close()

}
