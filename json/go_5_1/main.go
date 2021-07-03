package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	go_log "github.com/mmcomp/go-log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["msg"] = "Hello world"
	jsonResponse, _ := json.MarshalIndent(response, "", "")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonResponse))
	go_log.Log("New Req")
}

func main() {
	http.HandleFunc("/", handler)
	go_log.Log(http.ListenAndServe(":8080", nil))
}
