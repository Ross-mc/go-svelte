package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ross-mc/go-svelte-proj/server"
)

const PORT = "8080"

func main() {
	http.HandleFunc("/", server.Router)
	fmt.Printf("Listening for requests on port: %v\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
