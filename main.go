package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/niljimeno/citadel/modules"
	"github.com/niljimeno/citadel/web"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("searching for ", r.URL.Query())
		cmp := web.Search([]modules.Result{})
		cmp.Render(context.Background(), w)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
