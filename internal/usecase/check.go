package usecase

import (
	"os"
)

func CheckFile() bool {
	if _, err := os.Stat("../../internal/usecase/temp/tempcardapio"); err == nil {
		return true
	} else {
		return false
	}
}
