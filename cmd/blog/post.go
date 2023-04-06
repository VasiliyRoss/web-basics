package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type postPage struct {
	PostContent     []postPageData
}

type postPageData struct {
	Title        string  `db:"title"`
	Subtitle     string  `db:"subtitle"`
	ArticleImage string  `db:"post_image_url"`
	ArticleText  string  `db:"article_text"`
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
		postContent, err := postContent(db)
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
			PostContent:      postContent,
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

func postContent(db *sqlx.DB) ([]postPageData, error) {
	const query = `
		SELECT
			title, 
			subtitle, 
			post_image_url,
            article_text
		FROM
			post
		WHERE post_id = 1;
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var post []postPageData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&post, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return post, nil
}