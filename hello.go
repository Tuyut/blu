package main

import (
	"io"
	"log"
	"net/http"
)

const serverPort string = "8080"

var defaultNamespace map[string]string

func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong\n")
}

func get(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, defaultNamespace[r.URL.Query().Get("k")])
}

func set(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "SET")
}

func main() {
	// Initialize default namespace
	defaultNamespace = make(map[string]string)
	defaultNamespace["hello"] = "world"

	// Initialize http server
	http.HandleFunc("/v1/get", get)
	http.HandleFunc("/v1/set", set)
	http.HandleFunc("/ping", ping)
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}
