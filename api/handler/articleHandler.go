package handler

import (
	"encoding/json"
	"go-core/pkg/packages"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packages.Articles)
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packages.Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range packages.Articles {
		if article.Id == key {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(article)
		}
	}
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article packages.Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	packages.Articles = append(packages.Articles, article)

	json.NewEncoder(w).Encode(article)
}
