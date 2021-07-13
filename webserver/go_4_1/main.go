package main

import (
	"net/http"

	go_log "github.com/mmcomp/go-log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logger := go_log.Begin()
	defer logger.End()
	w.Write([]byte("Hello world"))
	logger.Log("New Req")
}

func main() {
	logger := go_log.Begin()
	defer logger.End()
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Error("Http Server Error happend", err)
	}
}
