package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title              string
	Subtitle           string
	Button		       string
	FeaturedPosts      []featuredPostData
	MostRecentPosts    []mostRecentPostsData	
}

type featuredPostData struct {
	Title       string
	Subtitle    string
	Author      string
	AuthorImg   string
	PublishDate string
}

type mostRecentPostsData struct {
	PostImage   string
	Title       string
	Subtitle    string
	Author      string
	AuthorImg   string
	PublishDate string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:         "Let's do it together.",
		Subtitle:	   "We travel the world in search of stories. Come along for the ride.",
		FeaturedPosts: featuredPosts(),
		MostRecentPosts: mostRecentPosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/post.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:         "Escape",
		FeaturedPosts: featuredPosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:    "The Road Ahead",
			Subtitle: "The road ahead might be paved - it might not be.",
			Author: "Mat Vogels",
			AuthorImg: "assets/images/avatars/mat_vogels.png",
			PublishDate: "September 25, 2015",
			
		},
		{
			Title:    "From Top Down",
			Subtitle: "Once a year, go someplace you’ve never been before.",
			Author: "William Wong",
			AuthorImg: "assets/images/avatars/william_wong.png",
			PublishDate: "September 25, 2015",
		}
	}
}

func mostRecentPosts() []mostRecentPostsData {
	return []mostRecentPostsData{
		{
			PostImage:
			Title:    "Still Standing Tall",
			Subtitle: "The road ahead might be paved - it might not be.",
			Author: "Mat Vogels",
			AuthorImg: "assets/images/avatars/mat_vogels.png",
			PublishDate: "September 25, 2015",
			
		},
		{
			PostImage: 
			Title:    "From Top Down",
			Subtitle: "Once a year, go someplace you’ve never been before.",
			Author: "William Wong",
			AuthorImg: "assets/images/avatars/william_wong.png",
			PublishDate: "September 25, 2015",
		},
	}
}