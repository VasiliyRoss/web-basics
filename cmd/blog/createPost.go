package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


type createPostRequest struct {
	Title   				string `json:"title"`
	Subtitle 				string `json:"subtitle"`  
  Author 	   			string `json:"author"`
  Author_url			string `json:"author_photo"`
  Publish_date    string `json:"publish_date"`
  Card_image_url  string `json:"post_image"`
  Post_image_url  string `json:"card_image"`
  Content					string `json:"content"`
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
    reqData, err := io.ReadAll(r.Body)
    if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		var req createPostRequest
		err =  json.Unmarshal(reqData, &req)
			if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		err = savePost(db, req)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Post was successfully published")
	}
}

func savePost(db *sqlx.DB, req createPostRequest) error {
	const query = `
	INSERT INTO
		post
	(
		title, 
		subtitle, 
		author, 
		author_url, 
		publish_date, 
		card_image_url, 
		post_image_url, 
		content
	)
	VALUES
	(
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?
	)
`

_, err := db.Exec(query, req.Title, req.Subtitle, req.Author, req.Author_url, req.Publish_date, req.Card_image_url, req.Post_image_url, req.Content) 
return err

}