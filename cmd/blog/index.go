package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	FeaturedPosts   []*featuredPostData
	MostRecentPosts []*mostRecentPostsData
}

type featuredPostData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	PostImage   string `db:"card_image_url"`
	Category    string `db:"category"`
	PostID      string `db:"post_id"`
	PostURL     string
}

type mostRecentPostsData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	PostImage   string `db:"card_image_url"`
	PostID      string `db:"post_id"`
	PostURL     string
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPosts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) 
			log.Println(err)
			return 
		}

		mostRecentPosts, err := mostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) 
			log.Println(err)
			return 
		}

		ts, err := template.ParseFiles("pages/index.html") 
		if err != nil {
			http.Error(w, "Internal Server Error", 500) 
			log.Println(err.Error())                    
			return                                      
		}

		data := indexPage{
			FeaturedPosts:   featuredPosts,
			MostRecentPosts: mostRecentPosts,
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

func featuredPosts(db *sqlx.DB) ([]*featuredPostData, error) {
	const query = `
		SELECT
			title, 
			subtitle, 
			author, 
			author_url, 
			publish_date, 
			card_image_url,
			post_id,
			COALESCE(category, '') AS category
		FROM
			post
		WHERE featured = 1
	`

	var posts []*featuredPostData 

	err := db.Select(&posts, query) 
	if err != nil {                 
		return nil, err
	}

	for _, posts := range posts {
		posts.PostURL = "/post/" + posts.PostID
	}

	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]*mostRecentPostsData, error) {
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
	`

	var posts []*mostRecentPostsData

	err := db.Select(&posts, query) 
	if err != nil {                 
		return nil, err
	}

	for _, posts := range posts {
		posts.PostURL = "/post/" + posts.PostID
	}

	return posts, nil
}