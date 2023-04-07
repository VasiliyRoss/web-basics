package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Импортируем для возможности подключения к MySQL
	"github.com/jmoiron/sqlx"
)

const (
	port = ":3000"
	startMessage = "Start server "
	dbDriverName = "mysql"
)

func main() {
	db, err := openDB() // Открываем соединение к базе данных в самом начале
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(db, dbDriverName) // Расширяем стандартный клиент к базе	

	mux := http.NewServeMux()
	mux.HandleFunc("/home", index(dbx))
	mux.HandleFunc("/1", post(dbx, 1))
	mux.HandleFunc("/2", post(dbx, 2))
	mux.HandleFunc("/3", post(dbx, 3))
	mux.HandleFunc("/4", post(dbx, 4))
	mux.HandleFunc("/5", post(dbx, 5))
	mux.HandleFunc("/6", post(dbx, 6))
	mux.HandleFunc("/7", post(dbx, 7))
	mux.HandleFunc("/8", post(dbx, 8))

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
	return sql.Open(dbDriverName, "root:12345Q@tcp(localhost:3306)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
}