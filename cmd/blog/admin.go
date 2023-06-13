package main

import (
	"encoding/base64"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type createPostRequest struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Author      string `json:"author"`
	AuthorImage string `json:"author_photo"`
	PublishDate string `json:"publish_date"`
	PostImage   string `json:"post_image"`
	CardImage   string `json:"card_image"`
	Content     string `json:"content"`
}

func admin(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("pages/admin.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	log.Println("Request completed successfully")
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		var req createPostRequest
		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		err = savePost(db, req)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
	postID := generateID(10)

	var authorImgUrl = "static/images/avatars/" + postID + ".png"
	var postImgUrl = "static/images/backgrounds/" + postID + ".png"
	var cardImgUrl = "static/images/cards/" + postID + ".png"

	err := saveImageToFile(cardImgUrl, req.CardImage)
	if err != nil {
		return err
	}

	err = saveImageToFile(postImgUrl, req.PostImage)
	if err != nil {
		return err
	}

	err = saveImageToFile(authorImgUrl, req.AuthorImage)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, req.Title, req.Subtitle, req.Author, "/"+authorImgUrl, req.PublishDate, "/"+cardImgUrl, "/"+postImgUrl, req.Content)
	return err
}

func saveImageToFile(filePath, base64String string) error {

	base64Image := strings.Split(base64String, ",")[1]

	img, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(img)
	if err != nil {
		return err
	}

	return nil
}

func generateID(length int) string {
	result := make([]byte, length)
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := 0; i < length; i++ {
		result[i] = chars[random.Intn(len(chars))]
	}
	return string(result)
}
