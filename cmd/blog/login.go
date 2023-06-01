package main

import (
	"html/template"
	"log"
	"net/http"
)

func loginPage(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("pages/login.html") 
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