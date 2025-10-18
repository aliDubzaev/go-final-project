package main

import (
	"log"

	"go1f/pkg/db"
	"go1f/pkg/server"
)

func main() {
	if err := db.Init("scheduler.db"); err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
