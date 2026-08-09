package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `CREATE TABLE scheduler (
id INTEGER PRIMARY KEY AUTOINCREMENT,
date CHAR(8) NOT NULL DEFAULT "",
title VARCHAR,
comment TEXT,
repeat VARCHAR);
CREATE INDEX idx_scheduler_date ON scheduler(date);
`

/*
|id|date|title|comment|repeat|
|--|----|-----|-------|------|
d 1
d 7
d 60
y
*/

var db *sql.DB

func Init(dbFile string) error {
	// 1. Проверяем существование файла
	_, err := os.Stat(dbFile)
	install := os.IsNotExist(err)

	// 2. Открываем базу данных
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	// Проверяем соединение с базой данных
	if err = db.Ping(); err != nil {
		return err
	}

	// 3. Если файл не существовал, выполняем SQL-команды из схемы
	if install {
		_, err = db.Exec(schema)
		if err != nil {
			db.Close() // Закрываем соединение в случае ошибки создания таблиц
			return err
		}
	}

	return nil
}
