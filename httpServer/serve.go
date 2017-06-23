package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	http.HandleFunc("/earth", handler2)

	http.HandleFunc("/test", handler3)

	log.Fatal(http.ListenAndServe(":3000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Last test\n")
}
