package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	os.Remove("/tmp/unix.sock")
	listener, err := net.Listen("unix", "/tmp/unix.sock")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// Create the client
	go func() {
		conn, err := net.Dial("unix", "/tmp/unix.sock")
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		for {
			time.Sleep(1 * time.Second)
			_, err := conn.Write([]byte("Hello, Unix Domain Socket!"))
			if err != nil {
				panic(err)
			}
		}
	}()

	// Listen to connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			buf := make([]byte, 1024)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					panic(err)
				}
				fmt.Printf("Received: %s\n", string(buf[:n]))
			}
		}(conn)
	}
}
