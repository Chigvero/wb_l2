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

	GetSite()
}

func GetSite() {
	req, err := http.Get("https://www.google.com/")
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
