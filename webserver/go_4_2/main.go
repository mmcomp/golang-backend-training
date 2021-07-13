package main

import (
	"fmt"
	"net/http"

	go_log "github.com/mmcomp/go-log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logger := go_log.Begin()
	defer logger.End()
	fmt.Fprint(w, "Hello world")
	logger.Log("New Req")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	logger := go_log.Begin()
	defer logger.End()
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		logger.Error("No Name Error Happend")
		return
	}
	fmt.Fprintf(w, "Hello %s", name)
	logger.Highlightf("New Req to hello name = %q", name)
}

func main() {
	logger := go_log.Begin()
	defer logger.End()
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error("Error from Web Server", err)
	}
}
