package main

import (
	"net/http"
	"html"
	"fmt"
	"log"
)

func main() {
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}

	http.HandleFunc("/", httpHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}