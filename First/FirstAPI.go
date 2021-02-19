package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, _ *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello world"},
	}
	fmt.Println("Endpoint Hit: All article endpoint")
	_ = json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Homepage endpoint hit")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
func main() {
	handleRequests()
}
