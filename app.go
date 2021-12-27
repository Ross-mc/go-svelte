package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const PORT = "8080"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	fmt.Printf("Listening for requests on port: %v\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
