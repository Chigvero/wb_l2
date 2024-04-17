package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url1 = ":8080"

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello w"))
}
func main() {
	fmt.Println(("Start"))

	http.HandleFunc("/", HttpHandler)
	err := http.ListenAndServe(url1, nil)
	if err != nil {
		panic(err)
	}
}

func GetSite() {
	req, err := http.Get("http://localhost:8080/a")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	file, err := os.Create("site")
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(file, req.Body)
}
