package main

import (
	"fmt"
	"go-core/api/handler"
	"go-core/pkg/packages"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", handler.Homepage)
	myRouter.HandleFunc("/all", handler.ReturnAllArticles)
	myRouter.HandleFunc("/article/{id}", handler.ReturnSingleArticle)
	myRouter.HandleFunc("/article", handler.CreateNewArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	packages.Articles = []packages.Article{
		packages.Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		packages.Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
