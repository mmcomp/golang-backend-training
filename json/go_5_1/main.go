package main

import (
	"encoding/json"
	"net/http"

	go_log "github.com/mmcomp/go-log"
)

func jsonError(w http.ResponseWriter, error []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, err := w.Write(error)
	if err != nil {
		go_log.Error("Write Json Error to output Error", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	go_log.Begin()
	defer go_log.End()
	response := map[string]string{}
	response["msg"] = "Hello world"
	jsonResponse, err := json.MarshalIndent(response, "", "")
	if err != nil {
		go_log.Error("Error Marshaling", err)
		response["msg"] = "Error in Marshaling output!"
		errorJsonResponse, _ := json.MarshalIndent(response, "", "")
		jsonError(w, errorJsonResponse, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, werr := w.Write(jsonResponse)
	if werr != nil {
		go_log.Error("Write to output Error", werr)
	}
	go_log.Log("New Req")
}

func main() {
	go_log.Begin()
	defer go_log.End()
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		go_log.Error("Http Error", err)
	}
	go_log.Log("Http server stopped!")
}
