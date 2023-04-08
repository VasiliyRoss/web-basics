package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type postPage struct {
	PostContent []postPageData
}

type postPageData struct {
	Title        string `db:"title"`
	Subtitle     string `db:"subtitle"`
	ArticleImage string `db:"post_image_url"`
	ArticleText  string `db:"article_text"`
}

func post(db *sqlx.DB, postId int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postContent, err := postContent(db, postId)
		if err != nil {
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

func postContent(db *sqlx.DB, postId int) ([]postPageData, error) {
	const query = `
		SELECT
			title, 
			subtitle, 
			COALESCE(post_image_url, '') AS post_image_url,
            COALESCE(article_text, '') AS article_text
		FROM
			post
		WHERE post_id = ?`

	var post []postPageData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&post, query, postId) // Делаем запрос в базу данных
	if err != nil {                        // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return post, nil
}
