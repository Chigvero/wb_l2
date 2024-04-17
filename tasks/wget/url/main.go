package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	parsed, err := url.Parse(
		"http://user:passw@93.184.216.34:8081/some/path/to/resource?val1=1&val2=2#anchor")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parsed.String())         // http://user:passw@93.184.216.34:8081/some/path/to/resource?val1=1&val2=2#anchor
	fmt.Println(parsed.User)             // user:passw
	fmt.Println(parsed.Scheme)           // http
	fmt.Println(parsed.Host)             // 93.184.216.34:8081
	fmt.Println(parsed.Path)             // /some/path/to/resource
	fmt.Println(parsed.Query().Encode()) // val1=1&val2=2
	fmt.Println(parsed.Fragment)         // anchor

	// worknig with elements directly
	parsed.Host = "myhost.com"
	parsed.Path = "new_path"
	parsed.RawQuery = "va3=three"
	parsed.Fragment = ""
	parsed.Scheme = "https"
	parsed.User = nil

	fmt.Println("------")

	fmt.Println(parsed.String())         // https://myhost.com/new_path?va3=three
	fmt.Println(parsed.User)             // ""
	fmt.Println(parsed.Scheme)           // https
	fmt.Println(parsed.Host)             // myhost.com
	fmt.Println(parsed.Path)             // "new_path"
	fmt.Println(parsed.Query().Encode()) // "va3=three"
	fmt.Println(parsed.Fragment)         // ""
}
