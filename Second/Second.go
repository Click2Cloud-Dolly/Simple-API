package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

func testPostArticles(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Test Post endpoint worked")
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "Homepage endpoint hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main() {
	handleRequests()
}
