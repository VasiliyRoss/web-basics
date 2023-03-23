package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title              string
	Subtitle           string
	Button		         string
	FeaturedPosts      []featuredPostData
	MostRecentPosts    []mostRecentPostsData	
}

type featuredPostData struct {
	PostImage   string
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
		Title:             "Let's do it together.",
		Subtitle:	         "We travel the world in search of stories. Come along for the ride.",
		Button:						 "View Latest Posts",
		FeaturedPosts:     featuredPosts(),
		MostRecentPosts:   mostRecentPosts(),
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
			PostImage: "background-image: url(assets/images/cards/the_road_ahead.png)",
			Title:    "The Road Ahead",
			Subtitle: "The road ahead might be paved - it might not be.",
			Author: "Mat Vogels",
			AuthorImg: "assets/images/avatars/mat_vogels.png",
			PublishDate: "September 25, 2015",	
		},
		{
			PostImage: "assets/images/cards/from_top_down.png",
			Title:    "From Top Down",
			Subtitle: "Once a year, go someplace you’ve never been before.",
			Author: "William Wong",
			AuthorImg: "assets/images/avatars/william_wong.png",
			PublishDate: "September 25, 2015",
		},
	}
}

func mostRecentPosts() []mostRecentPostsData {
	return []mostRecentPostsData{
		{
			PostImage: "assets/images/cards/still_standing_tall.jpg",
			Title:    "Still Standing Tall",
			Subtitle: "Life begins at the end of your comfort zone.",
			Author: "William Wong",
			AuthorImg: "assets/images/avatars/william_wong.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage: "assets/images/cards/sunny_side_up.png",
			Title:    "Sunny Side Up",
			Subtitle: "No place is ever as bad as they tell you it’s going to be",
			Author: "Mat Vogels",
			AuthorImg: "assets/images/avatars/mat_vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage: "assets/images/cards/water_falls.png",
			Title:    "Water Falls",
			Subtitle: "We travel not to escape life, but for life not to escape us.",
			Author: "Mat Vogels",
			AuthorImg: "assets/images/avatars/mat_vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage: "assets/images/cards/through_the_mist.png",
			Title:    "Through the Mist",
			Subtitle: "Travel makes you see what a tiny place you occupy in the	world.",
			Author: "William Wong",
			AuthorImg: "assets/images/avatars/william_wong.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage: "assets/images/cards/awaken_early.png",
			Title:    "Awaken Early",
			Subtitle: "Not all those who wander are lost.",
			Author: "Mat Vogels",
			AuthorImg: "assets/images/avatars/mat_vogels.png",
			PublishDate: "9/25/2015",
		},
		{
			PostImage: "assets/images/cards/try_it_always.png",
			Title:    "Try it Always",
			Subtitle: "The world is a book, and those who do not travel read only one page.",
			Author: "Mat Vogels",
			AuthorImg: "assets/images/avatars/mat_vogels.png",
			PublishDate: "9/25/2015",
		},
	}
}