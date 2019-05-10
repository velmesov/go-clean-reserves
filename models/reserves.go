package models

import (
	"GoProjects/clean-reserves/error"
	"database/sql"
	"fmt"
	"time"
)

// Clean чистит старые записи
func Clean(db *sql.DB) {
	// Подготавливаем запрос на удаление старых записей
	stmt, err := db.Prepare("DELETE FROM `reserves` WHERE `res_date_expiration` <= ? AND `res_unlimited` = ?")
	error.CheckNil(err)

	// Получаем текущую дату
	date := time.Now()

	// Переводим в строку и обрезаем сразу лишнее
	dateToString := fmt.Sprintf("%s", date)[0:19]

	_, err1 := stmt.Exec(dateToString, 0)
	error.CheckNil(err1)
}
