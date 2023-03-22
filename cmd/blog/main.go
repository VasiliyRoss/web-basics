package main

import (
	"log"
	"net/http"
)

const (
	port = ":3000"
	startMessage = "Start server "
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", index)
	mux.HandleFunc("/post", post)
	
	// Реализуем отдачу статики
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println(startMessage + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}