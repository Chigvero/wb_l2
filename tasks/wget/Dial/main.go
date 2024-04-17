package main

import (
	"fmt"
	"net"
)

func main() {
	// httpRequest := "GET / HTTP/1.1\n" +
	// 	"Host: golang.org\n\n"
	// conn, err := net.Dial("tcp", "golang.org:80")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer conn.Close()

	// if _, err = conn.Write([]byte(httpRequest)); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// io.Copy(os.Stdout, conn)
	// fmt.Println("Done")

	message := "Hello, I am a server" // отправляемое сообщение
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
