package main

import (
	"fmt"
	"net/http"

	go_log "github.com/mmcomp/go-log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	go_log.OutputFn("BEGIN")
	defer go_log.OutputFn("END")
	fmt.Fprint(w, "Hello world")
	go_log.Log("New Req")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	go_log.OutputFn("BEGIN")
	defer go_log.OutputFn("END")
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		go_log.Error("No Name Error Happend")
		return
	}
	fmt.Fprintf(w, "Hello %s", name)
	go_log.Logf("New Req to hello name = %q", name)
}

func main() {
	go_log.OutputFn("BEGIN")
	defer go_log.OutputFn("END")
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		go_log.Error("Error from Web Server", err)
	}
}
