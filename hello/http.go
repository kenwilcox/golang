package main

import(
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)

	addr := ":8080"
	fmt.Println("Starting web server on", addr)
	http.ListenAndServe(addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello %s\n", "World")
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello %s\n", "Earth")
}