package usecase

import (
	"fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateFile(msg string) {

	file, err := os.Create("../../internal/usecase/temp/tempcardapio")
	check(err)

	_, err = file.WriteString(msg)
	check(err)

	file.Close()

	go func() {
		fmt.Println("Timer On...")
		timer := time.NewTimer(30 * time.Minute)

		<-timer.C
		DeleteFile()
	}()

}

func DeleteFile() {
	err := os.Remove("../../internal/usecase/temp/tempcardapio")
	check(err)

}
