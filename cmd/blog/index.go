package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	FeaturedPosts      []featuredPostData
	MostRecentPosts    []mostRecentPostsData
}

type featuredPostData struct {
	Title       string  `db:"title"`
	Subtitle    string  `db:"subtitle"`
	Author      string  `db:"author"`
	AuthorImg   string  `db:"author_url"`
	PublishDate string  `db:"publish_date"`
	PostImage   string  `db:"card_image_url"`
	Category    string  `db:"category"`
	Id   		int     `db:"post_id"`
}

type mostRecentPostsData struct {
	Title       string  `db:"title"`
	Subtitle    string  `db:"subtitle"`
	Author      string  `db:"author"`
	AuthorImg   string  `db:"author_url"`
	PublishDate string  `db:"publish_date"`
	PostImage   string  `db:"card_image_url"`
	Id   		int     `db:"post_id"`
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
		featuredPosts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}
		
		mostRecentPosts, err := mostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

	    ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	    if err != nil {
		    http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		    log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		    return                                      // Не забываем завершить выполнение ф-ии
	    }

		data := indexPage{
			FeaturedPosts:      featuredPosts,
			MostRecentPosts:    mostRecentPosts,
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
			if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
		log.Println("Request completed successfully")
	}
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title, 
			subtitle, 
			author, 
			author_url, 
			publish_date, 
			card_image_url,
			post_id,
			COALESCE(category, '') AS category    /* чтобы не было ошибки в случае category = NULL */
		FROM
			post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []featuredPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]mostRecentPostsData, error) {
	const query = `
		SELECT
			title, 
			subtitle, 
			author, 
			author_url, 
			publish_date, 
			card_image_url,
			post_id
		FROM
			post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции featured-posts
	
	var posts []mostRecentPostsData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}