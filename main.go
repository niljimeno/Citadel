package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var page []byte

func main() {
	loadPage()
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadPage() error {
	f, err := os.Open("web/index.html")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	bufio.NewReader(f)
	page, err = io.ReadAll(f)
	return err
}
