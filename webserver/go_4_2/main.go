package main

import (
	"fmt"
	"net/http"

	go_log "github.com/mmcomp/go-log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
	go_log.Log("New Req")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query();
	name := query.Get("name")
	fmt.Fprintf(w, "Hello %s", name)
	go_log.Logf("New Req to hello name = %q", name)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	go_log.Log(http.ListenAndServe(":8080", nil))
}
