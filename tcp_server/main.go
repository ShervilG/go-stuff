package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("TCP server is listening on localhost:8080")
	defer func() {
		listener.Close()
		fmt.Println("TCP server has been closed")
	}()

	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	data := make([]byte, 2048)
	n, err := conn.Read(data)
	if err != nil {
		panic(err)
	}

	println("Received:", string(data[:n]))
}
