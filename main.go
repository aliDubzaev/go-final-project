package main

import (
	"log"

	"github.com/aliDubzaev/go-final-project/pkg/db"
	"github.com/aliDubzaev/go-final-project/pkg/server"
)

func main() {
	if err := db.Init("scheduler.db"); err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
