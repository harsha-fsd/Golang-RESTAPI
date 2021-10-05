package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var articles []Article = []Article{
	{1, "COVID pill 'cuts risk of death or hospitalisation by half,' says Merck - Euronews", "Drugmaker Merck has said that its experimental pill for people sick with COVID-19 reduced hospitalisations and deaths by half, in a potential leap forward in the global fight against the pandemic."},
	{2, "Canary Islands volcano increasingly aggressive as Spain's leader announces emergency funds - CNN", "The Cumbre Vieja volcano on the Spanish island of La Palma is now erupting even more aggressively after weeks of gushing lava, Spain's Instituto Geográfico Nacional (IGN) said Sunday."},
	{3, "Conservative conference: UK in period of adjustment after Brexit, says PM - BBC News", "Boris Johnson refuses to rule out supply problems at Christmas as fuel shortages continue."},
}
var id int = 3

func main() {
	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/all", getAllArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", getArticleById).Methods("GET")
	r.HandleFunc("/articles/add", createNewArticle).Methods("POST")
	r.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	r.HandleFunc("/", home).Methods("GET")
	log.Println("Server running on 127.0.0.1:9999")
	log.Fatal(http.ListenAndServe(":9999", r))

}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
	w.WriteHeader(http.StatusOK)
}

func getArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for _, article := range articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, fmt.Sprintf("No article by id %d exists", id), http.StatusNotFound)

}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	article := Article{}
	json.Unmarshal(reqBody, &article)
	id++
	article.Id = id
	articles = append(articles, article)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(article)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	deleteId, _ := strconv.Atoi(key)
	for index, article := range articles {
		if article.Id == deleteId {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}
