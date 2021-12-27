package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const PORT = "8080"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, "Hello World\n")
	case "POST":
		io.WriteString(w, "Post request reponse\n")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/", defaultHandler)
	fmt.Printf("Listening for requests on port: %v\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
