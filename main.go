package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is an anonymous API!")
	})

	http.ListenAndServe(":8080", nil)
}
