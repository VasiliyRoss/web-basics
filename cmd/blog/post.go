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

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid order id", 403)
			log.Println(err)
			return
		}

		postContent, err := postContent(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Post not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return 
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())                 
			return                                      
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

	var post []postPageData

	err := db.Select(&post, query, postID)
	if err != nil {                       
		return nil, err
	}

	return post, nil
}