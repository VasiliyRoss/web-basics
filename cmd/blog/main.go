package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

const (
	port         = ":3000"
	startMessage = "Start server "
	dbDriverName = "mysql"
	dbConnection = "root:12345Q@tcp(localhost:3306)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true"
)

func main() {

	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(db, dbDriverName)

	mux := mux.NewRouter()
	mux.HandleFunc("/home", index(dbx))

	mux.HandleFunc("/admin", admin)

	mux.HandleFunc("/login", loginPage)

	mux.HandleFunc("/api/post", createPost(dbx)).Methods(http.MethodPost)

	mux.HandleFunc("/post/{postID}", post(dbx))

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println(startMessage + port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() (*sql.DB, error) {
	return sql.Open(dbDriverName, dbConnection)
}
