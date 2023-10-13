package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloworld)
	fmt.Printf("Server started in port :8080")
	log.Fatal(http.ListenAndServe("8080", nil))
}
