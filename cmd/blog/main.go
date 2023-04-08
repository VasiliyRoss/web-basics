package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql" // Импортируем для возможности подключения к MySQL
	"github.com/jmoiron/sqlx"
)

const (
	port         = ":3000"
	startMessage = "Start server "
	dbDriverName = "mysql"
	dbConnection = "root:12345Q@tcp(localhost:3306)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true"
)

func main() {

	db, err := openDB() // Открываем соединение к базе данных в самом начале
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(db, dbDriverName) // Расширяем стандартный клиент к базе

	mux := http.NewServeMux()
	mux.HandleFunc("/home", index(dbx))

	count := 8 //тут нужно брать кол-во из бд

	// Выводим все страницы
	for count != 0 {
		mux.HandleFunc("/"+strconv.Itoa(count), post(dbx, count))
		count--
	}

	// Реализуем отдачу статики
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println(startMessage + port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}

}

func openDB() (*sql.DB, error) {
	// Здесь прописываем соединение к базе данных
	return sql.Open(dbDriverName, dbConnection)
}
