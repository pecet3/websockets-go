package main

import "net/http"

func main() {
	setupAPI()
}

func setupAPI() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.ListenAndServe(":8080", nil)

}
