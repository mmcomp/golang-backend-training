package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	go_log "github.com/mmcomp/go-log"
)

func jsonError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

func handler(w http.ResponseWriter, r *http.Request) {
	go_log.Begin()
	defer go_log.End();
	response := map[string]string{}
	response["msg"] = "Hello world"
	jsonResponse, err := json.MarshalIndent(response, "", "")
	if err != nil {
		go_log.Error("Error Marshaling", err)
		response["msg"] = "Error in Marshaling output!"
		errorJsonResponse, _ := json.MarshalIndent(response, "", "")
		jsonError(w, string(errorJsonResponse), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonResponse))
	go_log.Log("New Req")
}

func main() {
	go_log.Begin()
	defer go_log.End();
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		go_log.Error("Http Error", err)
	}
}
