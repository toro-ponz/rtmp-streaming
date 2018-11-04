package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	fmt.Println("Hello, world!")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
