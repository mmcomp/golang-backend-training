package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func additionHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)
	x, errx := strconv.ParseUint(query.Get("x"), 10, 64);
	if  errx != nil {
		response["error"] = errx.Error()
		jsonResponse, _ := json.MarshalIndent(response, "", "")
		fmt.Fprint(w, string(jsonResponse))
		go_log.Logf("%q is not a valid UINIT64", query.Get("x"))
		return
	}
	y, erry := strconv.ParseUint(query.Get("y"), 10, 64);
	if  erry != nil {
		response["error"] = erry.Error()
		jsonResponse, _ := json.MarshalIndent(response, "", "")
		fmt.Fprint(w, string(jsonResponse))
		go_log.Logf("%q is not a valid UINIT64", query.Get("y"))
		return
	}
	response["result"] = fmt.Sprint(x+y)
	jsonResponse, _ := json.MarshalIndent(response, "", "")
	fmt.Fprint(w, string(jsonResponse))
	go_log.Logf("New Req to addition %d + %d = %d", x, y, x+y)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/addition", additionHandler)
	go_log.Log(http.ListenAndServe(":8080", nil))
}
