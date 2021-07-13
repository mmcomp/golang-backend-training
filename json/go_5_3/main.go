package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	w.Write(jsonResponse)
	go_log.Log("New Req")
}

func additionHandler(w http.ResponseWriter, r *http.Request) {
	go_log.Begin()
	defer go_log.End()
	query := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{}
	var x uint64
	{
		var err error
	
		x, err = strconv.ParseUint(query.Get("x"), 10, 64)
		if  err != nil {
			response["error"] = err.Error()
			jsonResponse, _ := json.MarshalIndent(response, "", "")
			w.Write(jsonResponse)
			go_log.Logf("%q is not a valid UINIT64", query.Get("x"))
			return
		}
	}
	
	var y uint64
	{
		var err error
	
		y, err = strconv.ParseUint(query.Get("y"), 10, 64)
		if  err != nil {
			response["error"] = err.Error()
			jsonResponse, _ := json.MarshalIndent(response, "", "")
			w.Write(jsonResponse)
			go_log.Logf("%q is not a valid UINIT64", query.Get("y"))
			return
		}
	}
	response["result"] = fmt.Sprint(x+y)
	jsonResponse, _ := json.MarshalIndent(response, "", "")
	w.Write(jsonResponse)
	go_log.Logf("New Req to addition %d + %d = %d", x, y, x+y)
}

func main() {
	go_log.Begin()
	defer go_log.End()
	http.HandleFunc("/", handler)
	http.HandleFunc("/addition", additionHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		go_log.Error("Http Error", err)
	}
	go_log.Log("Http server stopped!")
}
