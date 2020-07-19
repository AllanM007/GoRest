package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article type to be used as database
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles global array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
