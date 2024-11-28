package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome!")

	fmt.Println("This is new line!")
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is 3rd anonymous API test!")
	})

	http.HandleFunc("/first", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is 3rd anonymous API test!")
	})

	fmt.Println("Added one new func!")
	http.HandleFunc("/anotherTest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is a another anonymous API!")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is an anonymous API!")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is a second anonymous API test!")
	})

	http.ListenAndServe(":8080", nil)
}
