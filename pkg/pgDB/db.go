package pgDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func Conn(
	host, user, password, dbname string,
	port int,
	sslmode bool,
) error {
	var ssl string
	// Проверяем sslmode
	if sslmode {
		ssl = "enable"
	} else {
		ssl = "disable"
	}
	// Получаем строку для подключения
	dataSourceName := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, ssl,
	)
	// Подключаемся к БД
	conn, errOpen := sql.Open("postgres", dataSourceName)

	if errOpen != nil {
		conn.Close()
		return errOpen
	}
	// Проверяем подключение к БД
	if err := conn.Ping(); err != nil {
		conn.Close()
		return err
	}

	DB = conn
	return nil
}
