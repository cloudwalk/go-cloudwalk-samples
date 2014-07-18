package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var addr *string = flag.String("h", "0.0.0.0:80", "address")
var delay *int = flag.Int("d", 0, "delay in secs to process POST")
var random *int = flag.Int("r", 0, "random delay range, -d must be 0 or not setted")

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "http://docs.cloudwalk.io", http.StatusFound)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if r.Method == "POST" {
		buf := r.URL.Query().Get("buf")
		anotherVariable := r.URL.Query().Get("anotherVariable")
		d := *delay
		if *delay == 0 && *random != 0 {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			d = r.Intn(*random)
		}
		time.Sleep(time.Duration(d) * time.Second)
		if anotherVariable != "" {
			fmt.Fprintf(w, "CLOUDWALK %s %s", buf, anotherVariable)
		} else {
			fmt.Fprintf(w, "CLOUDWALK %s", buf)
		}
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
