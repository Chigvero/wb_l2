package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8081"

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We're live !!!\n"))
}
func HomePageHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "We're live!!!\n")
	fmt.Fprintf(w, "We're live!!!\n")

}

func HomePageHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v, you are at %v. You have sent some query params too: %v.The method you used is %v", r.UserAgent(), r.URL.Path, r.URL.Query(), r.Method)

}
func main() {
	http.HandleFunc("/", HomePageHandler1)

	http.HandleFunc("/hi/", HomePageHandler)

	http.HandleFunc("/hello", HomePageHandler2)
	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
