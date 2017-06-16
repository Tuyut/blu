package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong\n")
}

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("Listening on http://localhost:8080")
	http.HandleFunc("/", hello)
	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
