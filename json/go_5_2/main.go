package main

import (
	"fmt"
	"net/http"
	"encoding/json"

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

func helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query();
	name := query.Get("name")
	response := make(map[string]string)
	response["msg"] = "Hello " + name
	jsonResponse, _ := json.MarshalIndent(response, "", "")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonResponse))
	go_log.Logf("New Req to hello name = %q", name)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	go_log.Log(http.ListenAndServe(":8080", nil))
}
