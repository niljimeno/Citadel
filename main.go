package main

import (
	"context"
	"log"
	"net/http"

	"github.com/niljimeno/citadel/repository"
	"github.com/niljimeno/citadel/web"
)

func main() {
	err := repository.Connect()
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search/", search)
	http.HandleFunc("/tag/", tag)

	log.Print("Running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	cmp := web.Search(repository.Search(query), "")
	cmp.Render(context.Background(), w)
}

func tag(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	cmp := web.Search(repository.FilterByTag(tag), "Filtering by "+tag)
	cmp.Render(context.Background(), w)
}
