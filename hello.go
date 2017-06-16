package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf("Listening on http://localhost:8080")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
