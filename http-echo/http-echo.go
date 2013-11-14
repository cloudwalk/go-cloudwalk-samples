package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "http://docs.cloudwalk.io", http.StatusFound)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if r.Method == "POST" {
		buf := r.URL.Query().Get("buf")
		anotherVariable := r.URL.Query().Get("anotherVariable")
		if anotherVariable != "" {
			fmt.Fprintf(w, "CLOUDWALK %s %s", buf, anotherVariable)
		} else {
			fmt.Fprintf(w, "CLOUDWALK %s", buf)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
