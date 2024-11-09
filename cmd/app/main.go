package main

import (
	"log"

	"github.com/jedyEvgeny/wallet-service/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal("ошибка на старте сервиса:", err)
	}
	err = a.Run()
	if err != nil {
		log.Fatal("ошибка выполнения сервиса:", err)
	}
}
