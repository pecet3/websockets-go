package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	setupAPI()
	log.Println("> Starting the server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI() {

	ctx := context.Background()

	manager := NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
}
