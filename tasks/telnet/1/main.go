package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("START")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	str = str[:len(str)-1]
	telStr := strings.Split(str, " ")
	if len(telStr) != 3 || telStr[0] != "telnet" {
		fmt.Println("Using: telnet host  port")
		return
	}
	httpRequest1, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	httpRequest2, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	httpRequest3, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// "GET / HTTP/1.1\n"
	//Host: golang.org\n\n
	httpRequest := httpRequest1 + httpRequest2 + httpRequest3
	conn, err := net.Dial("tcp", fmt.Sprint(telStr[1], ":", telStr[2]))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}
