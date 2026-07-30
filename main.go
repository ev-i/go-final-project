package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ev-i/go-final-project/pkg/db"
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

	// Определяем порт или используем порт по умолчанию
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	// Указываем директорию с фронтенд-файлами
	fs := http.FileServer(http.Dir("./web"))

	// Привязываем файловый сервер к корневому пути
	http.Handle("/", fs)

	log.Printf("Сервер запущен на http://localhost:%s\n", port)

	// Запускаем сервер
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}

}
