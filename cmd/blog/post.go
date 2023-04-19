package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type postPage struct {
	PostContent []postPageData
}

type postPageData struct {
	Title        string `db:"title"`
	Subtitle     string `db:"subtitle"`
	ArticleImage string `db:"post_image_url"`
	Content      string `db:"content"`
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"]

		postID, err := strconv.Atoi(postIDStr) // Конвертируем строку orderID в число
		if err != nil {
			http.Error(w, "Invalid order id", 403)
			log.Println(err)
			return
		}

		postContent, err := postContent(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				// sql.ErrNoRows возвращается, когда в запросе к базе не было ничего найдено
				// В таком случае мы возвращем 404 (not found) и пишем в тело, что ордер не найден
				http.Error(w, "Post not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
			return                                      // Не забываем завершить выполнение ф-ии
		}

		data := postPage{
			PostContent: postContent,
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		log.Println("Request completed successfully")
	}
}

func postContent(db *sqlx.DB, postID int) ([]postPageData, error) {
	const query = `
		SELECT
			title, 
			subtitle, 
			post_image_url,
            content
		FROM
			post
		WHERE post_id = ?`

	var post []postPageData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&post, query, postID) // Делаем запрос в базу данных
	if err != nil {                        // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return post, nil
}
