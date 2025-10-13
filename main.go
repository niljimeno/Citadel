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

	http.HandleFunc("/search/", func(w http.ResponseWriter, r *http.Request) {
		cmp := web.Search(repository.Search(r.URL.Query().Get("search")))
		cmp.Render(context.Background(), w)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
