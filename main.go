package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// this is for hello world get function
func helloGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "you are git another method please try agin with GET Resques", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := json.NewEncoder(w)
	response.Encode("hello world")
}
func main() {
	// this is for call hello get api
	app := http.NewServeMux()
	app.HandleFunc("/helloget", helloGet)

	fmt.Println("server run :3000")
	http.ListenAndServe(":3000", app)
}
