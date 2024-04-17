package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	req, err := http.Get("https://api.github.com/users")
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()
	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

}
