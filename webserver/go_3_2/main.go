package main

import (
	"fmt"
	"net/http"

	"github.com/mmcomp/go_log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world %s", r.URL.Query())
	go_log.Log("New Req")
}

func main() {
	http.HandleFunc("/", handler)
	go_log.Log(http.ListenAndServe(":8080", nil))
}
