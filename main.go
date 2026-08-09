package main

import (
	"log"
	"os"

	"github.com/ev-i/go-final-project/pkg/db"
	"github.com/ev-i/go-final-project/pkg/server"
)

func main() {

	// Определяем имя файла БД или используем имя по умолчанию
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	err := db.Init(dbFile)
	if err != nil {
		log.Fatal("Ошибка инициализации БД: ", err)
	}

	err = server.Run()
	if err != nil {
		log.Fatal("Ошибка запуска сервера", err)
	}

}
