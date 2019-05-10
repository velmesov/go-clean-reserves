package main

import (
	"GoProjects/clean-reserves/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Подключаемся к базе
	db := Db()

	// Закрываем соединение по завершению текущей функции
	defer db.Close()

	// Чистим старые резервы
	models.Clean(db)
}
