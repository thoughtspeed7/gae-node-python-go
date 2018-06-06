// om namah shivay

package main

import (
	"net/http"
	"log"
	"encoding/json"
)

func main() {
	
	http.HandleFunc("/", hello)
	
	log.Print("Go server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func hello(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	dictionary := make(map[string]string)
	dictionary["message"] = "Hello from Go!"

	json, _ := json.Marshal(dictionary)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
